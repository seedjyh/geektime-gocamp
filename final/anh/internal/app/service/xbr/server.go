package xbr

import (
	pb "anh/api"
	"context"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	err         error
	grpcAddress string
	grpcServer  *grpc.Server
}
type ServerOption func(s *Server)

func GRPCAddress(address string) ServerOption {
	return func(s *Server) {
		s.grpcAddress = address
	}
}

func NewServer(options ...ServerOption) (s *Server) {
	s = &Server{
		err:         nil,
		grpcAddress: "0.0.0.0:8082",
	}
	for _, opt := range options {
		opt(s)
	}
	// build up
	grpcServer := grpc.NewServer()
	pb.RegisterXBRServer(grpcServer, newXBRServer())
	s.grpcServer = grpcServer
	return s
}

// Start 启动并阻塞服务。
func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}

	lis, err := net.Listen("tcp", s.grpcAddress)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop(ctx context.Context) error {
	s.grpcServer.GracefulStop()
	return nil
}
