package ssh

import (
	"context"
	"time"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . SSHOptions
type SSHOptions interface {
	GetCredentials() (string, string)
	GetAddress() string
	GetTimeout() time.Duration
	GetCommand() string
}

//counterfeiter:generate . SSH
type SSH interface {
	Connect(options SSHOptions) error
	Run(context.Context, string, SSHRunData)
}

//counterfeiter:generate . SSHFactoryProvider
type SSHFactoryProvider interface {
	GetSSHFactory() func() SSH
}

//counterfeiter:generate . SSHRunData
type SSHRunData interface {
	StdoutChannel() chan []byte
	StderrChannel() chan []byte
	ErrorChannel() chan error
	DoneChannel() chan any
}
