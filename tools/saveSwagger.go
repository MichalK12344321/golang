package main

import (
	"context"
	"encoding/json"
	"flag"
	"lca/api/server"
	"lca/internal/app/cdg"
	"lca/internal/app/collector"
	"lca/internal/app/scheduler"
	"lca/internal/pkg/logging"
	"os"

	"github.com/defval/di"
	"github.com/zc2638/swag"
)

var outputPath *string
var appName *string

func main() {

	logger := logging.GetLogger()
	appName = flag.String("app", "", "app to get swagger from")
	outputPath = flag.String("out", "", "output path")
	flag.Parse()

	di.SetTracer(logger)

	_, err := di.New(
		getAppOptions(appName),
		di.Provide(server.NewContext),
		di.Provide(server.NewSwagApi),
		di.Provide(server.NewServeMux),
		di.Invoke(saveSwagger),
	)

	if err != nil {
		panic(err)
	}
}

func getAppOptions(app *string) di.Option {
	appName := *app
	switch appName {
	case "cdg":
		return cdg.ProvideOptions()
	case "scheduler":
		return scheduler.ProvideOptions()
	case "collector":
		return collector.ProvideOptions()
	default:
		panic("provide one of: cdg, scheduler, collector")
	}
}

func shutdown(ctx context.Context) {
	_, cancel := context.WithCancel(ctx)
	cancel()
}

func saveSwagger(ctx context.Context, api *swag.API, logger logging.Logger) error {
	defer shutdown(ctx)
	apiJson, err := json.MarshalIndent(api.Clone(), "", "  ")
	if err != nil {
		logger.Error("%s", err)
		return nil
	}
	path := *outputPath

	if path == "" {
		logger.Error("output path cannot be empty")
		return nil
	}

	err = os.WriteFile(path, apiJson, 0777)
	if err != nil {
		logger.Error("%s", err)
		return nil
	}
	return nil
}
