package main

import (
	appUser "anh/internal/app/interface/user"
	"anh/internal/pkg/app"
	"anh/internal/pkg/mylog"
	mySeelog "anh/internal/pkg/mylog/receiver/seelog"
	"flag"
	"github.com/BurntSushi/toml"
	"os"
)

var (
	appConfigFile = flag.String("app_config", "config.ini", "The config file for this application.")
	logConfigFile = flag.String("log_config", "seelog.xml", "The config file for the log")
)

const (
	appName    = "user-interface"
	appVersion = "1.0.0"
)

type web struct {
	Address string `toml:"address"`
}

type xbr struct {
	Address string `toml:"address"`
}

type config struct {
	Web web `toml:"web"`
	Xbr xbr `toml:"xbr"`
}

func main() {
	flag.Parse()
	mylog.Init(mySeelog.NewReceiverFromConfigAsFile(*logConfigFile))
	logger := mylog.CloneLogger().WithTag("app_name", appName)
	logger.Info("Start!")
	defer logger.Info("Exit!")
	var config config
	if _, err := toml.DecodeFile(*appConfigFile, &config); err != nil {
		_ = logger.CloneLogger().WithFields(
			mylog.Error("error", err), mylog.String("path", *appConfigFile)).
			Error("Invalid config")
		os.Exit(-1)
	}
	appServer := appUser.NewServer(
		appUser.WebAddress(config.Web.Address),
		appUser.XBRAddress(config.Xbr.Address),
	)
	a := app.New(
		appName,
		appVersion,
		appServer,
	)
	if err := a.Run(); err != nil {
		_ = logger.CloneLogger().WithFields(mylog.Error("error", err)).Error("Exit with error")
		os.Exit(-1)
	}
}
