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
	xbrAddress  string
	echoEngine  *echo.Echo
}
type ServerOption func(s *Server)

func WebAddress(address string) ServerOption {
	return func(s *Server) {
		s.webAddress = address
	}
}

func AuthAddress(address string) ServerOption {
	return func(s *Server) {
		s.authAddress = address
	}
}

func XBRAddress(address string) ServerOption {
	return func(s *Server) {
		s.xbrAddress = address
	}
}

func NewServer(options ...ServerOption) (s *Server) {
	s = &Server{
		err:         nil,
		webAddress:  "0.0.0.0:8080",
		authAddress: "127.0.0.1:8081",
		xbrAddress:  "127.0.0.1:8082",
	}
	for _, opt := range options {
		opt(s)
	}
	// build up
	xbrClient := newXBRClient(s.xbrAddress)
	bindingHandler := newBindingHandler(xbrClient)
	// start web service
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig)) // 允许跨站访问
	e.Pre(appendSessionID())
	e.Use(recordLatency(), dumpRequest(), dumpResponse())
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
