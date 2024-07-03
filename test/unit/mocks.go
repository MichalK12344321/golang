package test

import (
	"context"
	"lca/internal/app/cdg"
	"lca/internal/app/cdg/cdgfakes"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/broker/brokerfakes"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/ssh"
	"lca/internal/pkg/ssh/sshfakes"
	"lca/internal/pkg/storage"
	"lca/internal/pkg/storage/storagefakes"
	"lca/internal/pkg/util"

	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
)

type Mocks struct{}

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

func (m *Mocks) GetSSHFactoryProvider() (ssh.SSHFactoryProvider, *sshfakes.FakeSSH) {
	fakeSSH := &sshfakes.FakeSSH{}
	fakeProvider := &sshfakes.FakeSSHFactoryProvider{}
	factoryFn := func() ssh.SSH {
		return fakeSSH
	}
	fakeProvider.GetSSHFactoryReturns(factoryFn)
	return fakeProvider, fakeSSH
}

func (m *Mocks) GetCollectionScheduleEvent() *events.CollectionScheduleEvent {
	return &events.CollectionScheduleEvent{
		Id:     uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		Host:   "localhost",
		Port:   22,
		Script: "ls -la",
		SSHCredentials: events.SSHCredentials{
			User:     "user",
			Password: "password",
		},
	}
}

func (m *Mocks) GetCollectionJobUpdatedEvent() *events.CollectionJobUpdatedEvent {
	return &events.CollectionJobUpdatedEvent{
		CollectionScheduleEvent: *m.GetCollectionScheduleEvent(),
		Status:                  events.CollectionStatusInitialized,
		Error:                   "",
		Path:                    "",
	}
}
