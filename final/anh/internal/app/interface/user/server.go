package user

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	err         error
	webAddress  string
	authAddress string
	echoEngine  *echo.Echo
}

func NewServer(options ...ServerOption) (s *Server) {
	s = &Server{
		err:         nil,
		webAddress:  "0.0.0.0:8080",
		authAddress: "127.0.0.1:8081",
	}
	for _, opt := range options {
		opt(s)
	}
	// build up
	// authService := auth.New(...) 包含了gRPC调用auth
	// bindingService := bind.New(...) 包含了gRPC调用binding
	bindingHandler := &bindingHandler{}
	// start web service
	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/binding", bindingHandler.Bind)
	e.DELETE("/binding/:bind_id", bindingHandler.Unbind)
	s.echoEngine = e
	return s
}

// Start 启动并阻塞服务。
func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	return s.echoEngine.Start(s.webAddress)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.echoEngine.Shutdown(ctx)
}
