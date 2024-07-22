package storage

import (
	"path/filepath"
)

func GetStdOutPath(workDir string) string {
	return filepath.Join(workDir, "stdout.log")
}

func GetStdErrPath(workDir string) string {
	return filepath.Join(workDir, "stderr.log")
}
