package main

import "github.com/MenInBack/seelog"

func dispatcherFromFile() {
	logger, err := seelog.LoggerFromConfigAsFile("dispatcher-seelog.xml")
	if err != nil {
		panic(err)
	}
	defer logger.Flush()
	logger.Info("Hello, dispatcherFromFile")
}
