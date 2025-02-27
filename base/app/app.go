package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"qqlx/base/server"

	"go.uber.org/zap"
)

// Application is the main struct of the application
type Application struct {
	name    string
	version string
	servers []server.ServerInterface
	signals []os.Signal
}

// Option application support option
type Option func(application *Application)

// NewApp creates a new Application
func NewApp(ops ...Option) *Application {
	app := &Application{}
	for _, op := range ops {
		op(app)
	}

	// default accept signals
	if len(app.signals) == 0 {
		app.signals = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	}
	return app
}

// WithName application add name
func WithName(name string) func(application *Application) {
	return func(application *Application) {
		application.name = name
	}
}

// WithVersion application add version
func WithVersion(version string) func(application *Application) {
	return func(application *Application) {
		application.version = version
	}
}

// WithServer application add server
func WithServer(servers ...server.ServerInterface) func(application *Application) {
	return func(application *Application) {
		application.servers = servers
	}
}

// WithSignals application add listen signals
func WithSignals(signals []os.Signal) func(application *Application) {
	return func(application *Application) {
		application.signals = signals
	}
}

// Run application run
func (app *Application) Run(ctx context.Context) error {
	if len(app.servers) == 0 {
		return nil
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, app.signals...)
	errCh := make(chan error, 1)

	for _, s := range app.servers {
		go func(srv server.ServerInterface) {
			if err := srv.Start(); err != nil {
				zap.S().Errorf("failed to start server, err: %s", err)
				errCh <- err
			}
		}(s)
	}

	select {
	case err := <-errCh:
		_ = app.Stop()
		return err
	case <-ctx.Done():
		return app.Stop()
	case <-quit:
		return app.Stop()
	}
}

// Stop application stop
func (app *Application) Stop() error {
	wg := sync.WaitGroup{}
	for _, s := range app.servers {
		wg.Add(1)
		go func(srv server.ServerInterface) {
			defer wg.Done()
			if err := srv.Shutdown(); err != nil {
				zap.S().Errorf("failed to stop server, err: %s", err)
			}
		}(s)
	}
	wg.Wait()
	return nil
}
