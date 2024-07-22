package e2e_test

import (
	"context"
	"lca/internal/app/collector"
	"lca/internal/pkg/ssh"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	assert := assert.New(t)
	c := ssh.NewSSHClient()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := c.Connect(SSH_OPTIONS)
	assert.Nil(err)

	command := "pwd"

	runData := collector.NewCollectionJobData()

	go func() {
		c.Run(ctx, command, runData)
	}()

	for {
		select {
		case line := <-runData.StdoutChannel():
			assert.Equal("/root", line)
		case line := <-runData.StderrChannel():
			assert.Empty(line)
		case err := <-runData.ErrorChannel():
			assert.Nil(err)
		case <-runData.DoneChannel():
			return
		}
	}
}
