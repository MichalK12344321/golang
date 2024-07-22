package test

import (
	"lca/internal/pkg/events"

	"github.com/stretchr/testify/assert"
)

func AssertCollectionScheduleEvent(assert *assert.Assertions, expected *events.CollectionCreateEvent, actual *events.CollectionCreateEvent) {
	assert.Equal(expected.CollectionId, actual.CollectionId)
	if actual.Type == events.CollectionTypeSSH {
		assert.Equal(expected.SSH.Host, actual.SSH.Host)
		assert.Equal(expected.SSH.Port, actual.SSH.Port)
		assert.Equal(expected.SSH.Port, actual.SSH.Port)
		assert.Equal(expected.SSH.Script, actual.SSH.Script)
		assert.Equal(expected.SSH.User, actual.SSH.User)
		assert.Equal(expected.SSH.Password, actual.SSH.Password)
	}
}
