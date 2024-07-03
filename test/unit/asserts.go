package test

import (
	"lca/internal/pkg/events"

	"github.com/stretchr/testify/assert"
)

func AssertCollectionScheduleEvent(assert *assert.Assertions, expected *events.CollectionScheduleEvent, actual *events.CollectionScheduleEvent) {
	assert.Equal(expected.Id, actual.Id)
	assert.Equal(expected.Host, actual.Host)
	assert.Equal(expected.Port, actual.Port)
	assert.Equal(expected.Port, actual.Port)
	assert.Equal(expected.Script, actual.Script)
	assert.Equal(expected.User, actual.User)
	assert.Equal(expected.Password, actual.Password)
}
