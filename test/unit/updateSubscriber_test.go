package test

import (
	"context"
	"fmt"
	"lca/internal/app/cdg"
	"lca/internal/pkg/broker"
	"lca/internal/pkg/dto"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/util"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionUpdateSubscriberJob(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())

	repo, fakeRepo := mocks.GetLogCollectionRepository()

	expected := mocks.GetCollectionScheduleEvent()

	fakeRepo.InsertStub = func(actual *dto.CollectionDto) error {
		assert.Equal(expected.Id, actual.Id)
		assert.Equal(expected.Host, actual.Host)
		assert.Equal(expected.Port, actual.Port)
		assert.Equal(expected.Port, actual.Port)
		assert.Equal(expected.Script, actual.Script)
		assert.Equal(expected.User, actual.User)
		assert.Equal(expected.Password, actual.Password)
		return nil
	}

	payload, fakePayload := mocks.GetEventPayload(expected)
	fakePayload.AcknowledgeStub = func(err ...error) error {
		assert.False(util.AnyErrors(err...))
		assert.Equal(1, fakeRepo.UpdateCallCount())
		assert.Equal(1, fakePayload.AcknowledgeCallCount())
		cancel()
		waitGroup.Done()
		return nil
	}

	eventChannel := make(chan broker.EventPayload)
	go func() {
		cdg.CollectionUpdateSubscriberJob(ctx, eventChannel, logging.GetLogger(), repo)
	}()
	eventChannel <- payload

	waitGroup.Wait()
}

func TestUpdateFailure(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())

	repo, fakeRepo := mocks.GetLogCollectionRepository()

	expected := mocks.GetCollectionScheduleEvent()

	fakeRepo.UpdateStub = func(actual *dto.CollectionDto) error {
		return fmt.Errorf("Update failed")
	}

	payload, fakePayload := mocks.GetEventPayload(expected)
	fakePayload.AcknowledgeStub = func(err ...error) error {
		assert.True(util.AnyErrors(err...))
		assert.Equal(1, fakeRepo.UpdateCallCount())
		assert.Equal(1, fakePayload.AcknowledgeCallCount())
		cancel()
		waitGroup.Done()
		return nil
	}

	eventChannel := make(chan broker.EventPayload)
	go func() {
		cdg.CollectionUpdateSubscriberJob(ctx, eventChannel, logging.GetLogger(), repo)
	}()
	eventChannel <- payload

	waitGroup.Wait()
}
