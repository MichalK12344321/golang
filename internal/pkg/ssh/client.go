package ssh

import (
	"context"
	"sync"
	"time"

	"github.com/yahoo/vssh"
)

type SSHClient struct {
	vs      *vssh.VSSH
	timeout time.Duration
}

func NewSSHClient() *SSHClient {
	return &SSHClient{}
}

func (c *SSHClient) Connect(options SSHOptions) error {
	c.vs = vssh.New().Start()
	c.timeout = options.GetTimeout()
	config := vssh.GetConfigUserPass(options.GetCredentials())
	for _, addr := range []string{options.GetAddress()} {
		c.vs.AddClient(addr, config, vssh.SetMaxSessions(1))
	}
	_, err := c.vs.Wait()
	return err
}

func (c *SSHClient) Run(ctx context.Context, command string, runData SSHRunData) {
	respChan := c.vs.Run(ctx, command, c.timeout)

	var waitGroup sync.WaitGroup
	for resp := range respChan {
		if err := resp.Err(); err != nil {
			runData.ErrorChannel() <- err
			continue
		}

		waitGroup.Add(1)
		go func(resp *vssh.Response) {
			stream := resp.GetStream()
			defer stream.Close()

			for stream.ScanStdout() {
				runData.StdoutChannel() <- stream.BytesStdout()
			}

			for stream.ScanStderr() {
				runData.StderrChannel() <- stream.BytesStderr()
			}

			if err := stream.Err(); err != nil {
				runData.ErrorChannel() <- err
			}

			waitGroup.Done()
		}(resp)
	}

	waitGroup.Wait()
	close(runData.DoneChannel())
}

var _ SSH = new(SSHClient)
