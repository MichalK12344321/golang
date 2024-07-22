package test

import (
	"context"
	"lca/internal/app/collector"
	"lca/internal/pkg/events"
	"lca/internal/pkg/ssh"
	"lca/internal/pkg/util"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestManagerStart(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	createEvent := mocks.GetCollectionScheduleEvent()
	sshFactory, fakeSSH := mocks.GetSSHFactoryProvider()
	fakeSSH.RunStub = func(ctx context.Context, s string, sd ssh.SSHRunData) {
		assert.Equal(createEvent.SSH.Script, s)
		sd.StdoutChannel() <- []byte("first line")
		sd.StdoutChannel() <- []byte("second line")
		sd.DoneChannel() <- nil
	}
	_, fakeSink := mocks.GetSink()
	messageCount := 0
	expectedMessage := []string{"first line", "second line"}
	fakeSink.AppendStdoutStub = func(b []byte) {
		assert.Equal(string(b), expectedMessage[messageCount])
		messageCount++
	}
	sinkFactory, fakeSinkFactory := mocks.GetSinkFactory(fakeSink)
	brokerConnection, _ := mocks.GetBrokerConnection()
	publisherFactoryProvider, fakeEventPublisher := mocks.GetEventPublisherFactoryProvider()
	fakeEventPublisher.PublishStub = func(a any) error {
		runUpdate, ok := a.(*events.RunUpdateEvent)
		assert.Equal(createEvent.CollectionId, runUpdate.CollectionId)
		assert.True(ok)
		return nil
	}
	repo, fakeRepo := mocks.GetCollectorRepository()
	createEventJson, _ := util.ToJson(createEvent)
	fakeRepo.GetEventReturns(&collector.Event{
		EventId: createEvent.CollectionId,
		Type:    util.GetSimpleName(createEvent),
		Raw:     pgtype.JSONB{Bytes: createEventJson},
	}, nil)

	runner := collector.NewSSHRunner(sshFactory)
	publisher := collector.NewPublisher(ctx, brokerConnection, logger, publisherFactoryProvider)
	payload, fakePayload := mocks.GetEventPayload(mocks.GetRunCreateEvent())

	var manager collector.JobManager = collector.NewCollectionJobManager(
		ctx,
		logger,
		sshFactory,
		sinkFactory,
		publisher,
		repo,
		[]collector.CollectionRunner{runner},
	)

	var wg sync.WaitGroup
	wg.Add(1)

	fakeSink.FinalizeStub = func() error {
		wg.Done()
		return nil
	}

	err := manager.Start(payload)
	assert.Nil(err)
	wg.Wait()

	assert.Equal(2, messageCount)
	assert.Equal(1, fakePayload.AcknowledgeCallCount())
	assert.Equal(2, fakeEventPublisher.PublishCallCount())
	assert.Equal(1, fakeRepo.GetEventCallCount())
	assert.Equal(1, fakeRepo.SaveEventCallCount())
	assert.Equal(1, fakeSSH.ConnectCallCount())
	assert.Equal(1, fakeSSH.RunCallCount())
	assert.Equal(1, fakeSinkFactory.NewCallCount())
	assert.Equal(1, fakeSink.InitializeCallCount())
	assert.Equal(1, fakeSink.FinalizeCallCount())
	assert.Equal(2, fakeSink.AppendStdoutCallCount())
	assert.Equal(0, fakeSink.AppendStderrCallCount())
}

func TestManagerTerminate(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	createEvent := mocks.GetCollectionScheduleEvent()
	sshFactory, fakeSSH := mocks.GetSSHFactoryProvider()

	_, fakeSink := mocks.GetSink()
	sinkFactory, fakeSinkFactory := mocks.GetSinkFactory(fakeSink)
	brokerConnection, _ := mocks.GetBrokerConnection()
	publisherFactoryProvider, fakeEventPublisher := mocks.GetEventPublisherFactoryProvider()

	repo, fakeRepo := mocks.GetCollectorRepository()
	createEventJson, _ := util.ToJson(createEvent)
	fakeRepo.GetEventReturns(&collector.Event{
		EventId: createEvent.CollectionId,
		Type:    util.GetSimpleName(createEvent),
		Raw:     pgtype.JSONB{Bytes: createEventJson},
	}, nil)

	runner := collector.NewSSHRunner(sshFactory)
	publisher := collector.NewPublisher(ctx, brokerConnection, logger, publisherFactoryProvider)
	runEvent := mocks.GetRunCreateEvent()
	payload, fakePayload := mocks.GetEventPayload(runEvent)

	var manager collector.JobManager = collector.NewCollectionJobManager(
		ctx,
		logger,
		sshFactory,
		sinkFactory,
		publisher,
		repo,
		[]collector.CollectionRunner{runner},
	)

	var wg sync.WaitGroup
	wg.Add(2)

	fakeSSH.RunStub = func(ctx context.Context, s string, sd ssh.SSHRunData) {
		assert.Equal(createEvent.SSH.Script, s)
		wg.Done()
		<-ctx.Done()
	}

	fakeEventPublisher.PublishStub = func(a any) error {
		runUpdate, ok := a.(*events.RunUpdateEvent)
		assert.Equal(createEvent.CollectionId, runUpdate.CollectionId)
		assert.True(ok)
		if runUpdate.Status == events.CollectionStatusTerminated {
			wg.Done()
		}
		return nil
	}

	err := manager.Start(payload)
	assert.Nil(err)

	time.Sleep(time.Millisecond * 50)
	manager.Terminate(runEvent.RunId)

	wg.Wait()

	assert.Equal(1, fakePayload.AcknowledgeCallCount())
	assert.Equal(3, fakeEventPublisher.PublishCallCount())
	assert.Equal(1, fakeRepo.GetEventCallCount())
	assert.Equal(1, fakeRepo.SaveEventCallCount())
	assert.Equal(1, fakeSSH.ConnectCallCount())
	assert.Equal(1, fakeSSH.RunCallCount())
	assert.Equal(1, fakeSinkFactory.NewCallCount())
	assert.Equal(1, fakeSink.InitializeCallCount())
	assert.Equal(0, fakeSink.FinalizeCallCount())

	publishPayload, ok := fakeEventPublisher.PublishArgsForCall(0).(*events.RunUpdateEvent)
	assert.True(ok)
	assert.Equal(events.CollectionStatusStarted, publishPayload.Status)

	publishPayload, ok = fakeEventPublisher.PublishArgsForCall(1).(*events.RunUpdateEvent)
	assert.True(ok)
	assert.Equal(events.CollectionStatusTerminating, publishPayload.Status)

	publishPayload, ok = fakeEventPublisher.PublishArgsForCall(2).(*events.RunUpdateEvent)
	assert.True(ok)
	assert.Equal(events.CollectionStatusTerminated, publishPayload.Status)
}
