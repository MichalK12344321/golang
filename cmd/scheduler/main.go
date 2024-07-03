package main

import (
	"lca/api/server"
	"lca/internal/app/scheduler"
	"lca/internal/pkg/logging"
)

func main() {
	server.Run(logging.GetLogger(), scheduler.ProvideOptions())
}
