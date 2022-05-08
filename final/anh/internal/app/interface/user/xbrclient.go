package user

import (
	"context"
)

// XBRClient 是调用xbr服务，以实现绑定和解绑的客户端。
type XBRClient interface {
	Bind(ctx context.Context, parameter *BindParameter) (BindId, error)
	Unbind(ctx context.Context, bindId BindId) error
}
