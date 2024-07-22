package main

import (
	"lca/api/server"
	"lca/internal/app/cdg"
	"lca/internal/pkg/logging"
)

func main() {
	server.Run(logging.GetLogger(), cdg.ProvideOptions())
}
