package scheduler

import (
	"context"
	"lca/internal/config"
	"lca/internal/pkg/database"
	"lca/internal/pkg/dto"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"github.com/samber/lo"
)

type Scheduler struct {
	gocron.Scheduler
	logger logging.Logger
	config *config.Config

	publisher *Publisher

	dataContext database.DataContext
}

func NewScheduler(ctx context.Context, dataContext database.DataContext, logger logging.Logger, config *config.Config, createdPublisher *Publisher) (*Scheduler, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}
	result := &Scheduler{
		Scheduler:   scheduler,
		logger:      logger,
		config:      config,
		publisher:   createdPublisher,
		dataContext: dataContext,
	}

	go func() {
		<-ctx.Done()
		logger.Errors(result.Shutdown())
	}()

	go result.Start()

	return result, nil
}

func (s *Scheduler) PublishRunEvents(runEvents []*events.RunCreateEvent) error {
	for _, runEvent := range runEvents {
		err := s.publisher.PublishRunCreateEvent(runEvent)
		s.logger.Errors(err)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Scheduler) mapRunEvents(collectionIds []uuid.UUID) []*events.RunCreateEvent {
	runEvents := lo.Map(collectionIds, func(collectionId uuid.UUID, _ int) *events.RunCreateEvent {
		runEvent := &events.RunCreateEvent{
			RunId:        uuid.New(),
			CollectionId: collectionId,
		}
		return runEvent
	})
	return runEvents
}

func (s *Scheduler) ScheduleRuns(cronExpr string, repeat bool, collectionIds []uuid.UUID) ([]*events.RunCreateEvent, error) {
	runEvents := s.mapRunEvents(collectionIds)

	// one off immediately
	if cronExpr == "" {
		err := s.PublishRunEvents(runEvents)
		return runEvents, err
	}

	// cron job
	if repeat {
		job, err := s.Scheduler.NewJob(gocron.CronJob(cronExpr, false), gocron.NewTask(func() { s.PublishRunEvents(s.mapRunEvents(collectionIds)) }))
		s.logger.Errors(err)
		if err == nil {
			nextRun, _ := job.NextRun()
			s.logger.Info("job scheduled %d collection runs (periodic), next run at: %s", len(collectionIds), nextRun)
		}
		return make([]*events.RunCreateEvent, 0), err
	}

	// one off - at time
	parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	scheduleCron, err := parser.Parse(cronExpr)
	if err != nil {
		s.logger.Errors(err)
		return make([]*events.RunCreateEvent, 0), err
	}
	nextRun := scheduleCron.Next(time.Now())
	_, err = s.Scheduler.NewJob(
		gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(nextRun)), gocron.NewTask(func() { s.PublishRunEvents(s.mapRunEvents(collectionIds)) }),
	)
	s.logger.Errors(err)
	if err == nil {
		s.logger.Info("job scheduled %d collection runs, next run at: %s", len(collectionIds), nextRun)
	}
	return make([]*events.RunCreateEvent, 0), err
}

func (s *Scheduler) ScheduleSSH(schedule *dto.ScheduleSSHCollectionDto) (*dto.ScheduleResponseDto, error) {
	mappedEvents := schedule.ToEvents()

	s.logger.Errors(
		InsertEvent(s.dataContext, mappedEvents[0], mappedEvents[0].CollectionId),
		s.publisher.PublishCollectionCreateEvent(mappedEvents[0]),
	)

	runEvents, err := s.ScheduleRuns(schedule.Cron, schedule.Repeat, s.mapToId(mappedEvents))

	response := &dto.ScheduleResponseDto{
		Runs: lo.Map(runEvents, func(e *events.RunCreateEvent, _ int) *dto.RunScheduleDto { return (*dto.RunScheduleDto)(e) }),
	}

	return response, err
}

func (s *Scheduler) ScheduleGo(schedule *dto.ScheduleGoCollectionDto) (*dto.ScheduleResponseDto, error) {
	mappedEvents := schedule.ToEvents()

	s.logger.Errors(
		InsertEvent(s.dataContext, mappedEvents[0], mappedEvents[0].CollectionId),
		s.publisher.PublishCollectionCreateEvent(mappedEvents[0]),
	)

	runEvents, err := s.ScheduleRuns(schedule.Cron, schedule.Repeat, s.mapToId(mappedEvents))

	response := &dto.ScheduleResponseDto{
		Runs: lo.Map(runEvents, func(e *events.RunCreateEvent, _ int) *dto.RunScheduleDto { return (*dto.RunScheduleDto)(e) }),
	}

	return response, err
}

func (s *Scheduler) mapToId(evs []*events.CollectionCreateEvent) []uuid.UUID {
	return lo.Map(evs, func(e *events.CollectionCreateEvent, _ int) uuid.UUID { return e.CollectionId })
}

func (s *Scheduler) Terminate(t dto.TerminateRequestDto) error {
	terminateEvent := t.ToEvent()
	err := s.publisher.PublishRunTerminateEvent(terminateEvent)
	return err
}
