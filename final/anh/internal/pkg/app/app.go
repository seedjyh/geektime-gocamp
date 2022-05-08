// Package app 一个基本完整的应用生命周期管理包。

package app

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var stopped = errors.New("stopped")

type app struct {
	name        Name
	version     Version
	servers     []Server
	stopChannel chan struct{}
}

// New 创建一个 app。
func New(name Name, version Version, servers ...Server) *app {
	return &app{
		name:    name,
		version: version,
		servers: servers,
	}
}

// Run 启动所有 Service，并监听服务。
// 由于不是请求级别，所以不传入 context.Context
// 任何一个 Service 返回，或者收到 linux 信号，则全部停止。
func (a *app) Run() error {
	a.stopChannel = make(chan struct{}, 1)
	// 开始监听信号
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(signalChannel)
	// 协程生命周期控制
	var wg sync.WaitGroup
	defer wg.Wait()
	// 启动服务
	eg, egctx := errgroup.WithContext(context.Background())
	for _, s := range a.servers {
		wg.Add(1)
		s := s
		eg.Go(func() error {
			defer wg.Done()
			return s.Start(context.TODO())
		})
	}
	var err error
	select {
	case <-egctx.Done(): // 某个 errgroup 的 error 返回了
		a.stopAll()
		err = eg.Wait()
	case sig := <-signalChannel: // 收到了系统信号
		a.stopAll()
		err = &SignalCancel{sig}
	case <-a.stopChannel:
		a.stopAll()
		err = stopped
	}
	wg.Wait()
	return err
}

// Stop 停止所有 Service
func (a *app) Stop() {
	select {
	case a.stopChannel <- struct{}{}:
	default:
	}
}

func (a *app) stopAll() {
	for _, s := range a.servers {
		s.Stop(context.TODO())
	}
}
