package main

import (
	"errors"
	"test-seelog/mylog"
	"test-seelog/mylog/receiver/seelog"
)

func myJson() {
	mylog.Init(seelog.NewReceiverFromConfigAsFile("configs/my-seelog.xml"))
	logger := mylog.CloneLogger().
		WithTag("user_id", "user_123").     // tag 的值一定是字符串
		WithTag("session_id", "2147480001") // tag 的值一定是字符串
	_ = logger.WithFields(mylog.String("telX", "13336061916"), mylog.Int("status_code", 404), mylog.Float("cost_second", 0.1)).
		WithFields(mylog.Error("the_error", errors.New("some-error"))).
		Error("Hello, myJson!")
	logger.WithFields(mylog.String("state", "done")).
		Info("Bye, myJson!")
}
