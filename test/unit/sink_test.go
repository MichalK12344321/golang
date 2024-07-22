package test

import (
	"context"
	"lca/internal/app/collector"
	"lca/internal/pkg/events"
	"lca/internal/pkg/storage"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mholt/archiver/v3"
	"github.com/stretchr/testify/assert"
)

func TestStorageSink(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var jobData collector.JobData = collector.NewCollectionJobData()
	var sink collector.CollectionSink = collector.NewStorageSink(
		mocks.GetConfig(),
		ctx,
		mocks.GetRunCreateEvent(),
		logger,
		jobData,
	)
	populateSink(*assert, sink)

	filePath := filepath.Join(mocks.GetConfig().DataPath, mocks.GetRunCreateEvent().RunId.String()+".zip")
	defer os.Remove(filePath)
	extractedPath := strings.ReplaceAll(filePath, ".zip", "")
	err := archiver.NewZip().Unarchive(filePath, extractedPath)
	assert.Nil(err)

	stdoutContent, err := os.ReadFile(storage.GetStdOutPath(extractedPath))
	assert.Nil(err)
	assert.Equal("first line\nsecond line\n", string(stdoutContent))
	stderrContent, err := os.ReadFile(storage.GetStdErrPath(extractedPath))
	assert.Nil(err)
	assert.Equal("", string(stderrContent))
	assert.Nil(os.RemoveAll(extractedPath))
}

func TestBrokerSink(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	brokerConnection, _ := mocks.GetBrokerConnection()
	publisherFactoryProvider, fakeEventPublisher := mocks.GetEventPublisherFactoryProvider()

	var jobData collector.JobData = collector.NewCollectionJobData()
	event := mocks.GetRunCreateEvent()
	var sink collector.CollectionSink = collector.NewBrokerSink(
		ctx,
		event,
		brokerConnection,
		logger,
		jobData,
		publisherFactoryProvider,
	)
	populateSink(*assert, sink)

	assert.Equal(2, fakeEventPublisher.PublishCallCount())

	invocation := fakeEventPublisher.PublishArgsForCall(0).(*events.AppendLineEvent)
	assert.NotNil(invocation)
	assert.Equal("first line", invocation.Line)
	assert.Equal("stdout.log", invocation.File)
	assert.Equal(event.RunId, invocation.RunId)

	invocation = fakeEventPublisher.PublishArgsForCall(1).(*events.AppendLineEvent)
	assert.NotNil(invocation)
	assert.Equal("second line", invocation.Line)
	assert.Equal("stdout.log", invocation.File)
	assert.Equal(event.RunId, invocation.RunId)
}

func populateSink(assert assert.Assertions, sink collector.CollectionSink) {
	err := sink.Initialize()
	assert.Nil(err)

	sink.AppendStdout([]byte("first line"))
	sink.AppendStdout([]byte("second line"))

	err = sink.Finalize()
	assert.Nil(err)
}
