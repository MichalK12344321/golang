package test

import (
	"lca/internal/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CollectionJobChangedEvent struct{}

func TestGetSimpleName(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("CollectionJobChangedEvent", util.GetSimpleName(CollectionJobChangedEvent{}))
	assert.Equal("CollectionJobChangedEvent", util.GetSimpleName(&CollectionJobChangedEvent{}))
	var obj any = CollectionJobChangedEvent{}
	assert.Equal("CollectionJobChangedEvent", util.GetSimpleName(obj))
	assert.Equal("CollectionJobChangedEvent", util.GetSimpleName(&obj))
}
