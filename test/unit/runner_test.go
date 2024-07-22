package test

import (
	"context"
	"lca/internal/app/collector"
	"lca/internal/pkg/ssh"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoRunner(t *testing.T) {
	t.SkipNow()
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var waitGroup sync.WaitGroup

	var jobData collector.JobData = collector.NewCollectionJobData()
	var runner collector.CollectionRunner = collector.NewGoRunner(mocks.GetConfig(), logger)

	waitGroup.Add(1)
	go func() {
		assert.Equal([]byte("out\n"), <-jobData.StdoutChannel())
		waitGroup.Done()
	}()

	go runner.Run(ctx, jobData, `package main; import "fmt"; func main() {fmt.Println("out")}`)
	waitGroup.Wait()
}

func TestSSHRunner(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var waitGroup sync.WaitGroup

	sshFactory, fakeSSH := mocks.GetSSHFactoryProvider()
	sshOptions, fakeSSHOptions := mocks.GetSSHOptions()

	fakeSSHOptions.GetCommandReturns("echo out")

	fakeSSH.RunStub = func(ctx context.Context, s string, sd ssh.SSHRunData) {
		assert.Equal("echo out", s)
		sd.StdoutChannel() <- []byte("out\n")
		close(sd.DoneChannel())
	}

	var jobData collector.JobData = collector.NewCollectionJobData()
	var runner collector.CollectionRunner = collector.NewSSHRunner(sshFactory)

	waitGroup.Add(1)
	go func() {
		assert.Equal([]byte("out\n"), <-jobData.StdoutChannel())
		waitGroup.Done()
	}()

	go runner.Run(ctx, jobData, sshOptions)
	waitGroup.Wait()

	assert.Equal(1, fakeSSH.ConnectCallCount())
}
