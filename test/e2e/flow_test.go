package e2e_test

import (
	"context"
	"fmt"
	"io"
	"lca/internal/pkg/events"
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
	scheduleResponse, httpResSchedule, err := schedulerClient.ScheduleApi.PostSchedule(ctx, scheduler_client.LcaInternalPkgDtoScheduleRequestDto{
		Host:     "ne",
		Password: "root",
		User:     "root",
		Port:     22,
		Script:   "echo start && sleep 2 && echo end",
	})
	assert.Nil(err)
	assert.Equal(200, httpResSchedule.StatusCode)
	assert.NotEmpty(scheduleResponse.Id)

	time.Sleep(time.Millisecond * 300)

	cdgClientConfig := cdg_client.NewConfiguration()
	cdgClientConfig.BasePath = CDG_URL
	cdgClient := cdg_client.NewAPIClient(cdgClientConfig)
	collection, httpResCollection, err := cdgClient.CollectionApi.GetCollectionId(ctx, scheduleResponse.Id)
	assert.Nil(err)
	assert.Equal(string(events.CollectionStatusStarted), collection.Status)
	assert.Equal(200, httpResCollection.StatusCode)

	time.Sleep(time.Second * 3)

	collection, httpResCollection, err = cdgClient.CollectionApi.GetCollectionId(ctx, scheduleResponse.Id)
	assert.Nil(err)
	assert.Equal(string(events.CollectionStatusSuccess), collection.Status)
	assert.Equal(200, httpResCollection.StatusCode)

	getArchiveResponse, err := cdgClient.CollectionApi.GetCollectionIdArchive(ctx, scheduleResponse.Id)
	assert.Nil(err)
	assert.Equal(200, getArchiveResponse.StatusCode)
	assert.Contains(getArchiveResponse.Header.Get("Content-Disposition"), fmt.Sprintf("%s.zip", scheduleResponse.Id))

	tempPath, err := os.MkdirTemp(os.TempDir(), "")
	defer os.RemoveAll(tempPath)
	assert.Nil(err)

	filePath := filepath.Join(tempPath, fmt.Sprintf("%s.zip", scheduleResponse.Id))
	// TODO: using separate download because generated client does not return body but reads it
	err = downloadFile(filePath, fmt.Sprintf("%s/collection/%s/archive", CDG_URL, scheduleResponse.Id))
	assert.Nil(err)
	extractedPath := strings.ReplaceAll(filePath, ".zip", "")
	err = archiver.NewZip().Unarchive(filePath, extractedPath)
	assert.Nil(err)

	stdoutContent, err := os.ReadFile(filepath.Join(extractedPath, "stdout.log"))
	assert.Nil(err)
	assert.Equal("start\nend\n", string(stdoutContent))

	stderrContent, err := os.ReadFile(filepath.Join(extractedPath, "stderr.log"))
	assert.Nil(err)
	assert.Equal(``, string(stderrContent))
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
