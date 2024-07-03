package e2e_test

import (
	"context"
	"lca/internal/pkg/logging"
	"lca/internal/pkg/ssh"
	"lca/internal/pkg/ssh/sshfakes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	assert := assert.New(t)

	var c ssh.SSH = ssh.NewSSHClient()

	opts := &sshfakes.FakeSSHOptions{}
	opts.GetAddressReturns("localhost:30022")
	opts.GetCredentialsReturns("root", "root")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c.Connect(opts)

	command := "ls -la"

	stdOut, stdErr, err := c.Run(ctx, command)
	assert.Nil(err)
	assert.Empty(stdErr)
	assert.NotEmpty(stdOut)

	logging.GetLogger().Debug("%s", stdOut)
}
