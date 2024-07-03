package ssh

import (
	"context"
	"time"

	"github.com/yahoo/vssh"
)

type SSHClient struct {
	vs *vssh.VSSH
}

func NewSSHClient() *SSHClient {
	return &SSHClient{}
}

func (c *SSHClient) Disconnect() error {
	return nil
}

func (c *SSHClient) Connect(options SSHOptions) error {
	c.vs = vssh.New().Start()
	config := vssh.GetConfigUserPass(options.GetCredentials())
	for _, addr := range []string{options.GetAddress()} {
		c.vs.AddClient(addr, config, vssh.SetMaxSessions(1))
	}
	_, err := c.vs.Wait()
	return err
}

func (c *SSHClient) Run(ctx context.Context, command string) (string, string, error) {
	timeout, _ := time.ParseDuration("6s")
	respChan := c.vs.Run(ctx, command, timeout)

	for resp := range respChan {
		if err := resp.Err(); err != nil {
			return "", "", err
		}

		return resp.GetText(c.vs)
	}

	return "", "", nil
}

type VSSHFactoryProvider struct{}

func NewVSSHFactoryProvider() *VSSHFactoryProvider {
	return &VSSHFactoryProvider{}
}

func (p *VSSHFactoryProvider) GetSSHFactory() func() SSH {
	return func() SSH {
		return NewSSHClient()
	}
}

var _ SSHFactoryProvider = new(VSSHFactoryProvider)
