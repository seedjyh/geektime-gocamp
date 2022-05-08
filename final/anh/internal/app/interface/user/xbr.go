package user

import (
	pb "anh/api"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// xbr client
type xbrClient struct {
	address string
}

func newXBRClient(address string) *xbrClient {
	return &xbrClient{
		address: address,
	}
}

func (c *xbrClient) Bind(ctx context.Context, parameter *BindParameter) (BindId, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// connect
	conn, err := grpc.Dial(c.address, opts...)
	if err != nil {
		return "", errors.Wrap(err, "fail to dial")
	}
	defer conn.Close()
	client := pb.NewXBRClient(conn)
	// call
	if reply, err := client.Bind(ctx, &pb.BindRequest{
		TelA: parameter.TelA.String(),
		TelX: parameter.TelX.String(),
		TelB: parameter.TelB.String(),
	}); err != nil {
		return "", errors.Wrap(err, "fail to bind")
	} else {
		return BindId(reply.BindId), nil
	}
}

func (c *xbrClient) Unbind(ctx context.Context, bindId BindId) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// connect
	conn, err := grpc.Dial(c.address, opts...)
	if err != nil {
		return errors.Wrap(err, "fail to dial")
	}
	defer conn.Close()
	client := pb.NewXBRClient(conn)
	// call
	if _, err := client.Unbind(ctx, &pb.UnbindRequest{
		BindId: bindId.String(),
	}); err != nil {
		return errors.Wrap(err, "fail to unbind")
	} else {
		return nil
	}
}
