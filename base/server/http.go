package server

import (
	"context"
	"errors"
	"net/http"
	"time"
	"qqlx/base/conf"
	"qqlx/base/middleware"
	"qqlx/router"

	"github.com/gin-gonic/gin"
)

const DefaultShutdownTimeout = time.Second * 30

type Server struct {
	ShutdownTimeout time.Duration
	srv             *http.Server
}

type Options func(*Server)

func NewServer(e *gin.Engine, options ...Options) *Server {
	addr := conf.GetServerBind()
	ser := Server{
		ShutdownTimeout: DefaultShutdownTimeout,
		srv: &http.Server{
			Addr:    addr,
			Handler: e,
		},
	}

	for _, option := range options {
		option(&ser)
	}

	return &ser
}

// WithShutdownTimeout duration of graceful shutdown
func WithShutdownTimeout(duration time.Duration) Options {
	return func(server *Server) {
		server.ShutdownTimeout = duration
	}
}

// Start to start the server and wait for it to listen on the given address
func (s *Server) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown shuts down the server and close with graceful shutdown duration
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
}

func NewHttpServer(
	apiRouter *router.ApiRoute,
	authorization *middleware.AuthorizationMiddleware,
) *gin.Engine {
	if conf.GetLogLevel() == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })
	r.Use(middleware.RequestIDMiddleware(), middleware.CorssDomainMiddleware(), middleware.ZapMiddleware(), gin.Recovery())

	baseGroup := r.Group("/api/v1")
	apiRouter.RegisterApiUserRoute(baseGroup, authorization)
	apiRouter.RegisterApiRoleRoute(baseGroup, authorization)
	apiRouter.RegisterApiPolicyRoute(baseGroup, authorization)
	return r
}
