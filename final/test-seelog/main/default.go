package main

import "github.com/MenInBack/seelog"

func defaultLog() {
	// 默认时间是纳秒。
	defer seelog.Flush()
	seelog.Info("Hello, defaultLog")
}

func configFromFile() {
	logger, err := seelog.LoggerFromConfigAsFile("default-seelog.xml")
	if err != nil {
		panic(err)
	}
	defer logger.Flush()
	logger.Info("Hello, configFromFile")
}
