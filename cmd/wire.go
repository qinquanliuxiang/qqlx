//go:build wireinject
// +build wireinject

package cmd

import (
	"context"
	"qqlx/base/app"
	"qqlx/base/conf"
	"qqlx/base/middleware"
	"qqlx/base/server"
	"qqlx/controller"
	"qqlx/router"
	"qqlx/service"
	"qqlx/store"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

const (
	Version = "1.0.0"
)

func newApplication(e *gin.Engine) *app.Application {
	return app.NewApp(
		app.WithName(conf.GetProjectName()),
		app.WithVersion(Version),
		app.WithServer(server.NewServer(e)),
	)
}
func InitApplication(ctx context.Context, cabinModelFile string) (*app.Application, func(), error) {
	panic(wire.Build(
		server.NewHttpServer,
		store.ProviderStore,
		service.ProviderService,
		controller.ProviderContr,
		middleware.ProviderMiddleware,
		router.ProviderRouter,
		newApplication,
	))
}
