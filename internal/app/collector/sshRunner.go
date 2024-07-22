package collector

import (
	"context"
	"fmt"
	"lca/internal/pkg/ssh"
)

type SSHRunner struct {
	sshFactory ssh.SSHFactoryProvider
}

func NewSSHRunner(sshFactory ssh.SSHFactoryProvider) *SSHRunner {
	return &SSHRunner{sshFactory: sshFactory}
}

func (runner *SSHRunner) Run(ctx context.Context, jobData JobData, options CollectionRunnerOptions) {
	sshOptions, ok := options.(ssh.SSHOptions)
	if !ok {
		jobData.ErrorChannel() <- fmt.Errorf("provided option is not a valid type")
	}

	client := runner.sshFactory.GetSSHFactory()()
	err := client.Connect(sshOptions)
	if err != nil {
		jobData.ErrorChannel() <- err
	}

	client.Run(ctx, sshOptions.GetCommand(), jobData)
}

var _ CollectionRunner = new(SSHRunner)
