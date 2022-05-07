package main

import "github.com/MenInBack/seelog"

func formatFromFile() {
	logger, err := seelog.LoggerFromConfigAsFile("format-seelog.xml")
	if err != nil {
		panic(err)
	}
	defer logger.Flush()
	logger.Info("Hello, dispatcherFromFile")
}
