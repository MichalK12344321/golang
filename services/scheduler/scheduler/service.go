package scheduler

import (
  "fmt"

  "github.com/MichalK12344321/golang/services/common/types"
)

type dummyTriggerMessageQueue struct{
}

func (dmq *dummyTriggerMessageQueue)SendTrigger(id types.CollectionIdType , sshInfo types.SshInfoType) error {
  fmt.Println("SEND TRIGGER")
  return nil
}


type Service struct {
  config *Config
  server *Server
}

func NewService() *Service {
  config := NewConfig()
  server := NewServer(config, &dummyTriggerMessageQueue{})
  return &Service{config: config, server: server}
}

func (s *Service) Start() {
  s.server.Start()
}
