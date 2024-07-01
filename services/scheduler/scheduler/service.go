package scheduler

import (
  "fmt"

  "github.com/MichalK12344321/golang/services/common/types"
)

type dummyMessageQueue struct{
}

func (dmq *dummyMessageQueue)SendTrigger(id types.CollectionIdType , sshInfo types.SshInfoType) error {
  fmt.Println("SEND TRIGGER")
  return nil
}


type Service struct {
  server *Server
}

func NewService() *Service {
  server := NewServer(&dummyMessageQueue{})
  return &Service{server: server}
}

func (s *Service) Start() {
  s.server.Start()
}
