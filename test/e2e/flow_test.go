package e2e_test

import (
	"context"
	"fmt"
	"io"
	"lca/internal/pkg/events"
	"lca/internal/pkg/storage"
	cdg_client "lca/test/e2e/client/cdg/go_client"
	scheduler_client "lca/test/e2e/client/scheduler/go_client"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mholt/archiver/v3"
	"github.com/stretchr/testify/assert"
)

func TestFlow(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	schedulerClientConfig := scheduler_client.NewConfiguration()
	schedulerClientConfig.BasePath = SCHEDULER_URL
	schedulerClient := scheduler_client.NewAPIClient(schedulerClientConfig)
	scheduleResponse, httpResSchedule, err := schedulerClient.ScheduleApi.PostScheduleSsh(ctx, scheduler_client.LcaInternalPkgDtoScheduleSshCollectionDto{
		Cron:   "",
		Repeat: false,
		Targets: []scheduler_client.LcaInternalPkgDtoTargetDto{
			{
				Host:     NE_IP,
				User:     NE_USER,
				Password: NE_PASS,
				Port:     NE_PORT,
			},
		},
		Script: "echo start && sleep 2 && echo end",
	})
	assert.Nil(err)
	assert.Equal(200, httpResSchedule.StatusCode)
	assert.Equal(1, len(scheduleResponse.Runs))

	var collectionId string = scheduleResponse.Runs[0].CollectionId

	time.Sleep(time.Millisecond * 200)

	cdgClientConfig := cdg_client.NewConfiguration()
	cdgClientConfig.BasePath = CDG_URL
	cdgClient := cdg_client.NewAPIClient(cdgClientConfig)
	collection, httpResCollection, err := cdgClient.CollectionApi.GetCollectionId(ctx, collectionId)
	assert.Nil(err)
	runId := collection.Runs[0].RunId
	assert.Equal(string(events.CollectionStatusStarted), collection.Runs[0].Status)
	assert.Equal(200, httpResCollection.StatusCode)

	time.Sleep(time.Second * 3)

	collection, httpResCollection, err = cdgClient.CollectionApi.GetCollectionId(ctx, collectionId)
	assert.Nil(err)
	assert.Equal(string(events.CollectionStatusSuccess), collection.Runs[0].Status)
	assert.Equal(200, httpResCollection.StatusCode)

	getArchiveResponse, err := cdgClient.CollectionApi.GetCollectionRunRunIdArchive(ctx, runId)
	assert.Nil(err)
	assert.Equal(200, getArchiveResponse.StatusCode)
	assert.Contains(getArchiveResponse.Header.Get("Content-Disposition"), fmt.Sprintf("%s.zip", runId))

	tempPath, err := os.MkdirTemp(os.TempDir(), "")
	assert.Nil(err)
	err = os.MkdirAll(tempPath, 0777)
	// defer os.RemoveAll(tempPath)
	assert.Nil(err)

	filePath := filepath.Join(tempPath, fmt.Sprintf("%s.zip", runId))
	// TODO: using separate download because generated client does not return body but reads it
	err = downloadFile(filePath, fmt.Sprintf("%s/collection/run/%s/archive", CDG_URL, runId))
	assert.Nil(err)
	extractedPath := strings.ReplaceAll(filePath, ".zip", "")
	err = archiver.NewZip().Unarchive(filePath, extractedPath)
	assert.Nil(err)

	stdoutContent, err := os.ReadFile(storage.GetStdOutPath(extractedPath))
	assert.Nil(err)
	assert.Equal("start\nend\n", string(stdoutContent))

	stderrContent, err := os.ReadFile(storage.GetStdErrPath(extractedPath))
	assert.Nil(err)
	assert.Equal(``, string(stderrContent))
}
func TestTerminate(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	schedulerClientConfig := scheduler_client.NewConfiguration()
	schedulerClientConfig.BasePath = SCHEDULER_URL
	schedulerClient := scheduler_client.NewAPIClient(schedulerClientConfig)
	scheduleResponse, httpResSchedule, err := schedulerClient.ScheduleApi.PostScheduleSsh(ctx, scheduler_client.LcaInternalPkgDtoScheduleSshCollectionDto{
		Cron:   "",
		Repeat: false,
		Targets: []scheduler_client.LcaInternalPkgDtoTargetDto{
			{
				Host:     NE_IP,
				User:     NE_USER,
				Password: NE_PASS,
				Port:     NE_PORT,
			},
		},
		Script: SSH_INFINITE_SCRIPT,
	})
	assert.Nil(err)
	assert.Equal(200, httpResSchedule.StatusCode)
	assert.Equal(1, len(scheduleResponse.Runs))

	run := scheduleResponse.Runs[0]

	time.Sleep(time.Millisecond * 300)

	cdgClientConfig := cdg_client.NewConfiguration()
	cdgClientConfig.BasePath = CDG_URL
	cdgClient := cdg_client.NewAPIClient(cdgClientConfig)
	collection, httpResCollection, err := cdgClient.CollectionApi.GetCollectionId(ctx, run.CollectionId)
	assert.Nil(err)
	assert.Equal(string(events.CollectionStatusStarted), collection.Runs[0].Status)
	assert.Equal(200, httpResCollection.StatusCode)

	_, terminateHttpResponse, err := schedulerClient.ScheduleApi.PostTerminate(ctx, scheduler_client.LcaInternalPkgDtoTerminateRequestDto{RunId: run.RunId})
	assert.Nil(err)
	assert.Equal(200, terminateHttpResponse.StatusCode)

	time.Sleep(time.Second * 1)

	collection, httpResCollection, err = cdgClient.CollectionApi.GetCollectionId(ctx, run.CollectionId)
	assert.Nil(err)
	assert.Equal(string(events.CollectionStatusTerminated), collection.Runs[0].Status)
	assert.Equal(200, httpResCollection.StatusCode)
}

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
