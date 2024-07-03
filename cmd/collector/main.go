package main

import (
	"lca/api/server"
	"lca/internal/app/collector"
	"lca/internal/pkg/logging"
)

func main() {
	server.Run(logging.GetLogger(), collector.ProvideOptions())
}
