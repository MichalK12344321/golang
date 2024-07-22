package test

import (
	"context"
	"lca/internal/app/cdg"
	"lca/internal/app/cdg/cdgfakes"
	"lca/internal/app/collector"
	"lca/internal/app/collector/collectorfakes"
	"lca/internal/config"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/broker/brokerfakes"
	"lca/internal/pkg/database"
	"lca/internal/pkg/database/databasefakes"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/ssh"
	"lca/internal/pkg/ssh/sshfakes"
	"lca/internal/pkg/storage"
	"lca/internal/pkg/storage/storagefakes"
	"lca/internal/pkg/util"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
)

type Mocks struct{}

func (m *Mocks) GetConfig() *config.Config {
	dataPath, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	return &config.Config{
		App:  config.App{DataPath: dataPath},
		HTTP: config.HTTP{},
		RMQ:  config.RMQ{},
	}
}

func (m *Mocks) GetRMQDelivery() *rabbitmq.Delivery {
	delivery := &rabbitmq.Delivery{}
	delivery.CorrelationId = "321"
	delivery.Exchange = "ex"
	delivery.RoutingKey = ""
	return delivery
}

func (m *Mocks) GetEventPayload(body any) (broker.EventPayload, *brokerfakes.FakeEventPayload) {
	result := &brokerfakes.FakeEventPayload{}
	if body != nil {
		var bodyJson []byte
		bodyString, isString := body.(string)

		if isString {
			bodyJson = []byte(bodyString)
		} else {
			bodyJson, _ = util.ToJson(body)
		}

		result.ParseBodyStub = func(target any, pbf broker.ParseBodyFn) error {
			return events.FromJson(target, bodyJson, util.GetSimpleName(target))
		}
	}
	return result, result
}

func (m *Mocks) GetBrokerConnection() (broker.BrokerConnection, *brokerfakes.FakeBrokerConnection) {
	result := &brokerfakes.FakeBrokerConnection{}
	return result, result
}

func (m *Mocks) SetupBrokerConnection(setupFn func() broker.BrokerConnection) broker.BrokerConnection {
	return setupFn()
}

func (m *Mocks) GetEventPublisherFactoryProvider() (broker.PublisherFactoryProvider, *brokerfakes.FakeEventPublisher) {
	fakePub := &brokerfakes.FakeEventPublisher{}
	fake := &brokerfakes.FakePublisherFactoryProvider{}
	factoryFn := func(ctx context.Context, connection broker.BrokerConnection, logger logging.Logger, pubConfig *broker.PublisherConfig) broker.EventPublisher {
		return fakePub
	}
	fake.GetPublisherFactoryReturns(factoryFn)
	return fake, fakePub
}

func (m *Mocks) GetLogCollectionRepository() (cdg.LogCollectionRepository, *cdgfakes.FakeLogCollectionRepository) {
	result := &cdgfakes.FakeLogCollectionRepository{}
	return result, result
}

func (m *Mocks) GetStorage() (storage.Storage, *storagefakes.FakeStorage) {
	result := &storagefakes.FakeStorage{}
	return result, result
}

func (m *Mocks) GetSessionValue() (collector.JobData, *collectorfakes.FakeJobData) {
	stdoutChannel := make(chan []byte)
	stderrChannel := make(chan []byte)
	errorChannel := make(chan error, 1)
	doneChannel := make(chan any, 1)
	result := &collectorfakes.FakeJobData{}
	result.StdoutChannelReturns(stdoutChannel)
	result.StderrChannelReturns(stderrChannel)
	result.ErrorChannelReturns(errorChannel)
	result.DoneChannelReturns(doneChannel)
	return result, result
}

func (m *Mocks) GetSSHFactoryProvider() (ssh.SSHFactoryProvider, *sshfakes.FakeSSH) {
	fakeSSH := &sshfakes.FakeSSH{}
	fakeProvider := &sshfakes.FakeSSHFactoryProvider{}
	factoryFn := func() ssh.SSH {
		return fakeSSH
	}
	fakeProvider.GetSSHFactoryReturns(factoryFn)
	return fakeProvider, fakeSSH
}

func (m *Mocks) GetSSHOptions() (ssh.SSHOptions, *sshfakes.FakeSSHOptions) {
	result := &sshfakes.FakeSSHOptions{}
	return result, result
}

func (m *Mocks) GetJobManager() (collector.JobManager, *collectorfakes.FakeJobManager) {
	result := &collectorfakes.FakeJobManager{}
	return result, result
}

func (m *Mocks) GetJob() (collector.Job, *collectorfakes.FakeJob) {
	result := &collectorfakes.FakeJob{}
	return result, result
}

func (m *Mocks) GetCollectionSink() (collector.CollectionSink, *collectorfakes.FakeCollectionSink) {
	result := &collectorfakes.FakeCollectionSink{}
	return result, result
}

func (m *Mocks) GetCollectionScheduleEvent() *events.CollectionCreateEvent {
	return &events.CollectionCreateEvent{
		CollectionId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		Type:         events.CollectionTypeSSH,
		SSH: &events.SSHInfo{
			Host:     "localhost",
			Port:     22,
			Script:   "ls -la",
			User:     "user",
			Password: "password",
			Timeout:  time.Hour,
		},
	}
}

// func (m *Mocks) GetCollectionJobUpdatedEvent() *events.RunUpdateEvent {
// 	return &events.RunUpdateEvent{
// 		CollectionScheduleEvent: *m.GetCollectionScheduleEvent(),
// 		Status:                  events.CollectionStatusCreated,
// 		Error:                   "",
// 		Path:                    "",
// 	}
// }

func (m *Mocks) GetRunCreateEvent() *events.RunCreateEvent {
	return &events.RunCreateEvent{
		RunId:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		CollectionId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
	}
}

func (m *Mocks) GetSink() (collector.CollectionSink, *collectorfakes.FakeCollectionSink) {
	result := &collectorfakes.FakeCollectionSink{}
	return result, result
}

func (m *Mocks) GetSinkFactory(sinks ...collector.CollectionSink) (collector.CollectionSinkFactory, *collectorfakes.FakeCollectionSinkFactory) {
	result := &collectorfakes.FakeCollectionSinkFactory{}
	result.NewStub = func(ctx context.Context, rce *events.RunCreateEvent, jd collector.JobData) []collector.CollectionSink {
		return sinks
	}
	return result, result
}

func (m *Mocks) GetDataContext() (database.DataContext, *databasefakes.FakeDataContext) {
	result := &databasefakes.FakeDataContext{}
	return result, result
}

func (m *Mocks) GetRunner() (collector.CollectionRunner, *collectorfakes.FakeCollectionRunner) {
	result := &collectorfakes.FakeCollectionRunner{}
	return result, result
}

func (m *Mocks) GetCollectorRepository() (collector.Repository, *collectorfakes.FakeRepository) {
	result := &collectorfakes.FakeRepository{}
	return result, result
}
