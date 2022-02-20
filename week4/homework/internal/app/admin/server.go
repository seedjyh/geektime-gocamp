package admin

import (
	"context"
	"fmt"
	"geektime-gocamp/week4/homework/internal/student"
	stuEcho "geektime-gocamp/week4/homework/internal/student/http/echo"
	stuMySQL "geektime-gocamp/week4/homework/internal/student/repo/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	err            error
	address        string
	dataSourceName string
	echoEngine     *echo.Echo
}

func NewServer(options ...ServerOption) (s *Server) {
	s = &Server{
		err:            nil,
		address:        "0.0.0.0:8080",
		dataSourceName: "root:123456@tcp(127.0.0.1:3306)/test_db",
	}
	for _, opt := range options {
		opt(s)
	}
	// build up
	var mysqlEngine *xorm.Engine
	if engine, err := xorm.NewEngine("mysql", s.dataSourceName); err != nil {
		s.err = fmt.Errorf("invalid DSN, error=[%+v]", err)
		return
	} else if err := engine.Ping(); err != nil {
		s.err = fmt.Errorf("mySQL connect failed, error=[%+v]", err)
		return
	} else {
		mysqlEngine = engine
	}
	studentRepo := stuMySQL.New(mysqlEngine)
	studentService := student.New(studentRepo)
	studentRegisterHandler := stuEcho.NewRegister(studentService)
	studentUnregisterHandler := stuEcho.NewUnregister(studentService)
	// start web service
	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/students/add", studentRegisterHandler.Handle)
	e.DELETE("/students/:uid", studentUnregisterHandler.Handle)
	s.echoEngine = e
	return s
}

// Start 启动并阻塞服务。
func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	return s.echoEngine.Start(s.address)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.echoEngine.Shutdown(ctx)
}
