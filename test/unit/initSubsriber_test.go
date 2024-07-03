package test

import (
	"context"
	"fmt"
	"lca/internal/app/collector"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/events"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/ssh"
	"lca/internal/pkg/util"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCollectionInitSubscriberJob(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())

	conn, _ := mocks.GetBrokerConnection()
	factory, fakePub := mocks.GetEventPublisherFactoryProvider()
	sshFactory, fakeSSH := mocks.GetSSHFactoryProvider()
	storage, fakeStorage := mocks.GetStorage()

	pub := collector.NewCollectionUpdatePublisher(ctx, conn, logger, factory)

	expected := mocks.GetCollectionScheduleEvent()

	fakeSSH.ConnectStub = func(options ssh.SSHOptions) error {
		user, pass := options.GetCredentials()
		assert.Equal(expected.User, user)
		assert.Equal(expected.Password, pass)
		assert.Equal(fmt.Sprintf("%s:%d", expected.Host, expected.Port), options.GetAddress())
		return nil
	}

	fakeSSH.RunStub = func(ctx context.Context, s string) (string, string, error) {
		assert.Equal(expected.Script, s)
		return "stdout\n", "stderr\n", nil
	}

	fakeStorage.CreateFileStub = func(id uuid.UUID, stdout, stderr string) (string, error) {
		assert.Equal(expected.Id, id)
		assert.Equal("stdout\n", stdout)
		assert.Equal("stderr\n", stderr)
		return "", nil
	}

	assertFns := []func(*events.CollectionJobUpdatedEvent){
		func(event *events.CollectionJobUpdatedEvent) {
			assert.Equal(events.CollectionStatusStarted, event.Status)
		},
		func(event *events.CollectionJobUpdatedEvent) {
			assert.Equal(events.CollectionStatusSuccess, event.Status)
		},
	}
	publishCalls := 0
	fakePub.PublishStub = func(a any) error {
		event, ok := a.(*events.CollectionJobUpdatedEvent)
		assert.True(ok)
		AssertCollectionScheduleEvent(assert, expected, &event.CollectionScheduleEvent)
		assertFns[publishCalls](event)
		publishCalls++
		return nil
	}

	payload, fakePayload := mocks.GetEventPayload(expected)
	fakePayload.AcknowledgeStub = func(err ...error) error {
		assert.False(util.AnyErrors(err...))
		assert.Equal(2, fakePub.PublishCallCount())
		assert.Equal(1, fakeStorage.CreateFileCallCount())
		assert.Equal(1, fakePayload.AcknowledgeCallCount())
		assert.Equal(1, fakeSSH.RunCallCount())
		cancel()
		waitGroup.Done()
		return nil
	}

	eventChannel := make(chan broker.EventPayload)
	go func() {
		collector.CollectionInitSubscriberJob(ctx, eventChannel, logging.GetLogger(), pub, storage, sshFactory)
	}()
	eventChannel <- payload

	waitGroup.Wait()
}

func TestParseFailure(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())

	conn, _ := mocks.GetBrokerConnection()
	factory, fakePub := mocks.GetEventPublisherFactoryProvider()
	sshFactory, fakeSSH := mocks.GetSSHFactoryProvider()
	storage, fakeStorage := mocks.GetStorage()

	pub := collector.NewCollectionUpdatePublisher(ctx, conn, logger, factory)

	bodyJson := `{ "id": false }` // id is a uuid

	payload, fakePayload := mocks.GetEventPayload(bodyJson)
	fakePayload.AcknowledgeStub = func(err ...error) error {
		assert.True(util.AnyErrors(err...))
		assert.Equal(1, fakePub.PublishCallCount())
		assert.Equal(0, fakeStorage.CreateFileCallCount())
		assert.Equal(1, fakePayload.AcknowledgeCallCount())
		assert.Equal(0, fakeSSH.RunCallCount())
		cancel()
		waitGroup.Done()
		return nil
	}

	eventChannel := make(chan broker.EventPayload)
	go func() {
		collector.CollectionInitSubscriberJob(ctx, eventChannel, logging.GetLogger(), pub, storage, sshFactory)
	}()
	eventChannel <- payload

	waitGroup.Wait()
}

func TestSSHFailure(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())

	conn, _ := mocks.GetBrokerConnection()
	factory, fakePub := mocks.GetEventPublisherFactoryProvider()
	sshFactory, fakeSSH := mocks.GetSSHFactoryProvider()
	storage, fakeStorage := mocks.GetStorage()

	pub := collector.NewCollectionUpdatePublisher(ctx, conn, logger, factory)

	expected := mocks.GetCollectionScheduleEvent()

	fakeSSH.ConnectStub = func(options ssh.SSHOptions) error {
		return fmt.Errorf("SSH fail")
	}

	payload, fakePayload := mocks.GetEventPayload(expected)
	fakePayload.AcknowledgeStub = func(err ...error) error {
		assert.True(util.AnyErrors(err...))
		assert.Equal(2, fakePub.PublishCallCount())
		assert.Equal(0, fakeStorage.CreateFileCallCount())
		assert.Equal(1, fakePayload.AcknowledgeCallCount())
		assert.Equal(0, fakeSSH.RunCallCount())
		cancel()
		waitGroup.Done()
		return nil
	}

	eventChannel := make(chan broker.EventPayload)
	go func() {
		collector.CollectionInitSubscriberJob(ctx, eventChannel, logging.GetLogger(), pub, storage, sshFactory)
	}()
	eventChannel <- payload

	waitGroup.Wait()
}

func TestSSHRunFailure(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())

	conn, _ := mocks.GetBrokerConnection()
	factory, fakePub := mocks.GetEventPublisherFactoryProvider()
	sshFactory, fakeSSH := mocks.GetSSHFactoryProvider()
	storage, fakeStorage := mocks.GetStorage()

	pub := collector.NewCollectionUpdatePublisher(ctx, conn, logger, factory)

	expected := mocks.GetCollectionScheduleEvent()

	fakeSSH.RunStub = func(ctx context.Context, s string) (string, string, error) {
		return "", "", fmt.Errorf("SSH run fail")
	}

	payload, fakePayload := mocks.GetEventPayload(expected)
	fakePayload.AcknowledgeStub = func(err ...error) error {
		assert.True(util.AnyErrors(err...))
		assert.Equal(2, fakePub.PublishCallCount())
		assert.Equal(0, fakeStorage.CreateFileCallCount())
		assert.Equal(1, fakePayload.AcknowledgeCallCount())
		assert.Equal(1, fakeSSH.RunCallCount())
		cancel()
		waitGroup.Done()
		return nil
	}

	eventChannel := make(chan broker.EventPayload)
	go func() {
		collector.CollectionInitSubscriberJob(ctx, eventChannel, logging.GetLogger(), pub, storage, sshFactory)
	}()
	eventChannel <- payload

	waitGroup.Wait()
}

func TestStorageFailure(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())

	conn, _ := mocks.GetBrokerConnection()
	factory, fakePub := mocks.GetEventPublisherFactoryProvider()
	sshFactory, fakeSSH := mocks.GetSSHFactoryProvider()
	storage, fakeStorage := mocks.GetStorage()

	pub := collector.NewCollectionUpdatePublisher(ctx, conn, logger, factory)

	expected := mocks.GetCollectionScheduleEvent()

	fakeStorage.CreateFileStub = func(u uuid.UUID, s1, s2 string) (string, error) {
		return "", fmt.Errorf("storage fail")
	}

	payload, fakePayload := mocks.GetEventPayload(expected)
	fakePayload.AcknowledgeStub = func(err ...error) error {
		assert.True(util.AnyErrors(err...))
		assert.Equal(2, fakePub.PublishCallCount())
		assert.Equal(1, fakeStorage.CreateFileCallCount())
		assert.Equal(1, fakePayload.AcknowledgeCallCount())
		assert.Equal(1, fakeSSH.RunCallCount())
		cancel()
		waitGroup.Done()
		return nil
	}

	eventChannel := make(chan broker.EventPayload)
	go func() {
		collector.CollectionInitSubscriberJob(ctx, eventChannel, logging.GetLogger(), pub, storage, sshFactory)
	}()
	eventChannel <- payload

	waitGroup.Wait()
}
