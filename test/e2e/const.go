package e2e_test

import (
	"lca/internal/pkg/ssh/sshfakes"
	"time"
)

const SCHEDULER_URL = "http://lca.localhost:8080/api/scheduler"
const CDG_URL = "http://lca.localhost:8080/api/cdg"
const COLLECTOR_URL = "http://lca.localhost:8080/api/collector"
const NE_IP = "192.168.49.2"
const NE_PORT = 30022
const NE_USER = "root"
const NE_PASS = "root"

const SSH_INFINITE_SCRIPT = "while :; do echo $(date); sleep 1; done"

var SSH_OPTIONS = &sshfakes.FakeSSHOptions{
	GetAddressStub:     func() string { return NE_IP },
	GetCredentialsStub: func() (string, string) { return NE_USER, NE_PASS },
	GetTimeoutStub:     func() time.Duration { return time.Duration(time.Hour * 1) },
}
