package ssh

import (
	"context"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . SSHOptions
type SSHOptions interface {
	GetCredentials() (string, string)
	GetAddress() string
}

//counterfeiter:generate . SSH
type SSH interface {
	Connect(options SSHOptions) error
	Disconnect() error
	Run(ctx context.Context, command string) (string, string, error)
}

//counterfeiter:generate . SSHFactoryProvider
type SSHFactoryProvider interface {
	GetSSHFactory() func() SSH
}
