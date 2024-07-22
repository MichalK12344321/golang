package collector

import (
	"context"
	"errors"
	"fmt"
	"lca/internal/config"
	"lca/internal/pkg/logging"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GoRunner struct {
	config      *config.Config
	programPath string
	logger      logging.Logger
}

func NewGoRunner(config *config.Config, logger logging.Logger) *GoRunner {
	cmd := exec.Command("whereis", "go")
	if cmd.Err != nil {
		panic(cmd.Err)
	}
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	paths := strings.ReplaceAll(strings.Split(string(output), ":")[1], "\n", "")
	programPath, ok := lo.Find(
		strings.Split(paths, " "),
		func(p string) bool {
			fi, err := os.Stat(p)
			exists := !errors.Is(err, os.ErrNotExist)
			return exists && !fi.IsDir() && fi.Mode()&0111 != 0
		},
	)

	if programPath == "" || !ok {
		panic("could not find go program path")
	}

	return &GoRunner{
		config:      config,
		programPath: programPath,
		logger:      logger,
	}
}

func (runner *GoRunner) AppendToFile(path string, content []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	return err
}

func (runner *GoRunner) Run(ctx context.Context, jobData JobData, options CollectionRunnerOptions) {
	script, ok := options.(string)
	if !ok {
		jobData.ErrorChannel() <- fmt.Errorf("provided option is not a valid type")
	}

	workDir := filepath.Join("/tmp", uuid.NewString())
	err := os.MkdirAll(workDir, 0777)
	if err != nil {
		jobData.ErrorChannel() <- err
	}

	scriptPath := filepath.Join(workDir, "main.go")
	err = runner.AppendToFile(scriptPath, []byte(script))
	if err != nil {
		jobData.ErrorChannel() <- err
	}

	cmd := exec.Command(runner.programPath, "run", scriptPath)
	if cmd.Err != nil {
		jobData.ErrorChannel() <- cmd.Err
	}

	output, err := cmd.Output()
	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok {
			jobData.StderrChannel() <- exitError.Stderr
		}
		jobData.ErrorChannel() <- err
	}
	jobData.StdoutChannel() <- output

	close(jobData.DoneChannel())
}

var _ CollectionRunner = new(GoRunner)
