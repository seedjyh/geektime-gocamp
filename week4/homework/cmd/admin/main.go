package main

import (
	"fmt"
	"geektime-gocamp/week4/homework/internal/students"
	stuEcho "geektime-gocamp/week4/homework/internal/students/http/echo"
	stuMySQL "geektime-gocamp/week4/homework/internal/students/repo/mysql"
	"github.com/BurntSushi/toml"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	configPath := "./week4/homework/cmd/admin/config.ini"
	var config config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		fmt.Printf("Invalid config, path=[%+v], error=[%+v]\n", configPath, err)
		os.Exit(-1)
	}
	// build up
	var mysqlEngine *xorm.Engine
	if engine, err := xorm.NewEngine("mysql", config.MySQL.DataSourceName); err != nil {
		fmt.Println("Invalid DSN, error=", err)
		os.Exit(-1)
	} else if err := engine.Ping(); err != nil {
		fmt.Println("MySQL connect failed, error=", err)
		os.Exit(-1)
	} else {
		mysqlEngine = engine
	}
	studentRepo := stuMySQL.New(mysqlEngine)
	studentService := students.New(studentRepo)
	studentRegisterHandler := stuEcho.NewRegister(studentService)
	studentUnregisterHandler := stuEcho.NewUnregister(studentService)
	// start web service
	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/students/add", studentRegisterHandler.Handle)
	e.DELETE("/students/:uid", studentUnregisterHandler.Handle)
	if err := e.Start(config.Web.Address); err != nil {
		fmt.Println("Start service failed, error=", err)
		os.Exit(-1)
	}
}
