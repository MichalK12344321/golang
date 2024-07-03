package server

import (
	"lca/internal/pkg/logging"
	"path"
	"strconv"

	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/defval/di"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/option"
)

func Run(logger logging.Logger, services di.Option) {
	di.SetTracer(logger)

	container, err := di.New(
		services,
		di.Provide(NewContext),
		di.Provide(NewServer),
		di.Provide(NewSwagApi),
		di.Provide(NewServeMux),
	)

	if err != nil {
		logger.Fatalf("%s", err)
	}

	if err := container.Invoke(StartServer); err != nil {
		logger.Fatalf("%s", err)
	}
}

func StartServer(ctx context.Context, server *http.Server, logger logging.Logger) error {
	logger.Info("Starting server... at %s", server.Addr)
	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()
	select {
	case <-ctx.Done():
		logger.Info("Stopping server...")
		return server.Close()
	case err := <-errChan:
		return fmt.Errorf("server error: %s", err)
	}
}

func NewContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
		<-stop
		cancel()
	}()
	return ctx
}

func NewServer(mux *http.ServeMux, serverConfig *ServerConfig) *http.Server {
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(serverConfig.Port),
		Handler: mux,
	}
	return server
}

func NewSwagApi(controllers []Controller, serverConfig *ServerConfig) *swag.API {
	api := swag.New(
		option.Title(serverConfig.Name),
		option.Version(serverConfig.Version),
		option.Description(""),
		// option.License("", ""),
		option.TermsOfService(""),
	)
	for _, controller := range controllers {
		api.AddEndpoint(controller.GetEndpoints()...)
	}
	return api
}

func NewServeMux(api *swag.API) *http.ServeMux {
	mux := &http.ServeMux{}
	for p, endpoints := range api.Paths {
		mux.Handle(path.Join(api.BasePath, p), endpoints)
	}
	mux.Handle("/api-docs", api.Handler())
	patterns := swag.UIPatterns("/api")

	for _, pattern := range patterns {
		mux.Handle(pattern, swag.UIHandler("/api", "/api-docs", true))
	}

	return mux
}
