# Week 3 Homework

Go进阶训练营（第7期）第3周作业。

## 概述

预先注册若干服务。

程序启动后，依次：

1. 监视系统信号`SIGTERM`和`SIGINT`。
1. 逐个启动注册的服务，用`waitgroup`确保协程退出，用`errgroup`获得首个返回的`error`。
1. 用`select`监视系统信号、`errgroup`以及一个`stopChannel`，最后那个是外部主动停止服务时使用。
