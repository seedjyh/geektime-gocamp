package app

import "context"

type Server interface {
	// Start 启动 Server 并阻塞，直到 Server 工作结束才返回。
	Start(ctx context.Context) error
	// Stop 通知 Server 停止工作。不阻塞等待。
	Stop(ctx context.Context) error
}
