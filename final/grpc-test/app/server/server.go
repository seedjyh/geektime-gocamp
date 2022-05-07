package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpc-test/api"
	"net"
)

type authServer struct {
	pb.UnimplementedAuthServer
	userid2token map[string]string
}

func newAuthServer() *authServer {
	return &authServer{
		userid2token: make(map[string]string),
	}
}

func (s *authServer) SetUserToken(ctx context.Context, idAndToken *pb.UserIdAndToken) (*emptypb.Empty, error) {
	fmt.Println("Setting user token:", idAndToken.UserId.Id, idAndToken.Token)
	s.userid2token[idAndToken.UserId.Id] = idAndToken.Token
	return &emptypb.Empty{}, nil
}

func (s *authServer) GetUserToken(ctx context.Context, userId *pb.UserId) (*pb.UserIdAndToken, error) {
	fmt.Println("Getting token for user:", userId.Id)
	if token, ok := s.userid2token[userId.Id]; ok {
		return &pb.UserIdAndToken{
			UserId: userId,
			Token:  token,
		}, nil
	} else {
		return nil, errors.New("not found")
	}
}

func Run(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, newAuthServer())
	return grpcServer.Serve(lis)
}
