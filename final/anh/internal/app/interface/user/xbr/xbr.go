// Package xbr 用gRPC方式实现了XBRClient接口，通过gRPC调用xbr服务，以实现绑定和解绑。
package xbr

import (
	pb "anh/api"
	"anh/internal/app/interface/user"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type client struct {
	address string
}

func NewClient(address string) *client {
	return &client{
		address: address,
	}
}

func (c *client) Bind(ctx context.Context, parameter *user.BindParameter) (user.BindId, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// connect
	conn, err := grpc.Dial(c.address, opts...)
	if err != nil {
		return "", errors.Wrap(err, "fail to dial")
	}
	defer conn.Close()
	client := pb.NewXBRClient(conn)
	mdCtx := metadata.AppendToOutgoingContext(ctx, "session_id", ctx.Value("session_id").(string))
	// call
	if reply, err := client.Bind(mdCtx, &pb.BindRequest{
		TelA: parameter.TelA.String(),
		TelX: parameter.TelX.String(),
		TelB: parameter.TelB.String(),
	}); err != nil {
		return "", errors.Wrap(err, "fail to bind")
	} else {
		return user.BindId(reply.BindId), nil
	}
}

func (c *client) Unbind(ctx context.Context, bindId user.BindId) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// connect
	conn, err := grpc.Dial(c.address, opts...)
	if err != nil {
		return errors.Wrap(err, "fail to dial")
	}
	defer conn.Close()
	client := pb.NewXBRClient(conn)
	mdCtx := metadata.AppendToOutgoingContext(ctx, "session_id", ctx.Value("session_id").(string))
	// call
	if _, err := client.Unbind(mdCtx, &pb.UnbindRequest{
		BindId: bindId.String(),
	}); err != nil {
		return errors.Wrap(err, "fail to unbind")
	} else {
		return nil
	}
}
