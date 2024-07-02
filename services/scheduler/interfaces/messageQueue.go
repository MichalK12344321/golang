package interfaces

import (
  "github.com/MichalK12344321/golang/services/common/types"
)

type TriggerMessageQueue interface {
  SendTrigger(types.CollectionIdType, types.SshInfoType) error
}
