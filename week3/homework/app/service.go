package app

type Service interface {
	// Run 在当前协程启动服务。如果已经启动，则直接失败。
	// 不是请求级别，所以不应该传入 context.Context
	Run() error
	// Stop 停止一个服务。如果已经停止，也不算错。
	Stop()
}
