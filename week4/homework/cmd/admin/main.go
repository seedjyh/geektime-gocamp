package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	configPath := "./week4/homework/cmd/admin/config.ini"
	var config config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		fmt.Println("Invalid config, path=", configPath, ", error=", err)
		os.Exit(-1)
	}
	e := echo.New()
	if err := e.Start(config.Web.Address); err != nil {
		fmt.Println("Start service failed, error=", err)
		os.Exit(-1)
	}
}
