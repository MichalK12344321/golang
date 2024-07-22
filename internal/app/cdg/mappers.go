package cdg

import (
	"lca/internal/pkg/dto"
	"lca/internal/pkg/events"

	"github.com/samber/lo"
)

func mapCollectionFromDto(c *Collection, _ int) *dto.CollectionDto {
	result := &dto.CollectionDto{
		CollectionCreateEvent: events.CollectionCreateEvent{
			CollectionId: c.Id,
			Type:         events.CollectionType(c.Type),
		},
	}
	if c.SSH != nil {
		result.SSH = &events.SSHInfo{
			User:     c.SSH.User,
			Password: c.SSH.Password,
			Host:     c.SSH.Host,
			Port:     c.SSH.Port,
			Script:   c.SSH.Script,
			Timeout:  c.SSH.Timeout,
		}
	}
	if c.Go != nil {
		result.Go = &events.GoInfo{
			Script:  c.Go.Script,
			Timeout: c.Go.Timeout,
		}
	}
	if len(c.Runs) > 0 {
		result.Runs = lo.Map(c.Runs, mapCollectionRunFromModel)
	} else {
		result.Runs = make([]*dto.RunDto, 0)
	}
	return result
}

func mapCollectionToModel(c *dto.CollectionDto) *Collection {
	result := &Collection{
		Id:   c.CollectionId,
		Type: c.Type,
	}

	if c.SSH != nil {
		result.SSH = &SshInfo{
			CollectionId: c.CollectionId,
			User:         c.SSH.User,
			Password:     c.SSH.Password,
			Host:         c.SSH.Host,
			Port:         c.SSH.Port,
			Script:       c.SSH.Script,
			Timeout:      c.SSH.Timeout,
		}
	}

	if c.Go != nil {
		result.Go = &GoInfo{
			CollectionId: c.CollectionId,
			Script:       c.Go.Script,
			Timeout:      c.Go.Timeout,
		}
	}

	return result
}

func mapCollectionRunToModel(c *dto.RunDto) *Run {
	result := &Run{
		Id:           c.RunId,
		CollectionId: c.CollectionId,
		Status:       c.Status,
		Error:        c.Error,
	}

	return result
}

func mapCollectionRunFromModel(c *Run, _ int) *dto.RunDto {
	result := &dto.RunDto{
		RunCreateEvent: events.RunCreateEvent{
			RunId:        c.Id,
			CollectionId: c.CollectionId,
		},
		Status: c.Status,
		Error:  c.Error,
	}

	return result
}
