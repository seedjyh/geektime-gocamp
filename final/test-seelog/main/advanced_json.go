package main

import (
	"github.com/MenInBack/seelog"
	"sync"
	"time"
)

type myLogContext struct {
	ReqId     string
	UserId    string
	SessionId string
}

func createReqIdFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		if myContext, ok := context.CustomContext().(*myLogContext); !ok {
			return "Broken Context!"
		} else {
			return myContext.ReqId
		}
	}
}

func advancedJsonFromFile() {
	if err := seelog.RegisterCustomFormatter("ReqId", createReqIdFormatter); err != nil {
		panic(err)
	}
	logger, err := seelog.LoggerFromConfigAsFile("configs/advanced-seelog.xml")
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger, err := seelog.CloneLogger(logger)
		if err != nil {
			panic(err)
		}
		defer logger.Flush()
		logger.SetContext(&myLogContext{ReqId: "odd-rid-001"})
		time.Sleep(time.Second)
		logger.Info("Hello, advancedJsonFromFile 1")
		time.Sleep(time.Second)
		logger.Info("Hello, advancedJsonFromFile 3")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger, err := seelog.CloneLogger(logger)
		if err != nil {
			panic(err)
		}
		defer logger.Flush()
		logger.SetContext(&myLogContext{ReqId: "even-rid-002"})
		time.Sleep(time.Second)
		logger.Info("Hello, advancedJsonFromFile 2")
		time.Sleep(time.Second)
		logger.Info("Hello, advancedJsonFromFile 4")
	}()
	wg.Wait()
}
