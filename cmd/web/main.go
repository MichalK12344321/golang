package main

import (
	"lca/api/server"
	"lca/internal/app/web"
	"lca/internal/pkg/logging"
)

func main() {
	server.Run(logging.GetLogger(), web.ProvideOptions())
}
