package test

import (
	"lca/internal/config"
	"lca/internal/pkg/storage"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	assert := assert.New(t)
	r, err := filepath.Abs(".")
	assert.NotNil(r)
	assert.Nil(err)

	var store storage.Storage = storage.NewDiskStorage(&config.Config{
		App:  config.App{DataPath: r},
		HTTP: config.HTTP{},
		RMQ:  config.RMQ{},
	})

	id := uuid.New()
	path, err := store.CreateFile(id, "std", "err")
	_ = path
	defer os.Remove(path)
	assert.Nil(err)

	files, err := store.ListFiles(id)
	assert.Nil(err)
	assert.Contains(files, "stdout.log")
}
