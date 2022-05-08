package xbr

import (
	pb "anh/api"
	"anh/internal/pkg/mylog"
	"context"
	"math/rand"
)

type xbrServer struct {
	pb.UnimplementedXBRServer
	id2data map[string]interface{}
}

func newXBRServer() *xbrServer {
	return &xbrServer{
		id2data: make(map[string]interface{}),
	}
}

func (s *xbrServer) Bind(ctx context.Context, request *pb.BindRequest) (*pb.BindReply, error) {
	telA := request.TelA
	telX := request.TelX
	telB := request.TelB
	mylog.CloneLogger().WithTag("app", "xbr-service").
		WithFields(mylog.String("tel_a", telA)).
		WithFields(mylog.String("tel_x", telX)).
		WithFields(mylog.String("tel_b", telB)).
		Info("received bind request")
	bindId := s.randomBindID()
	mylog.CloneLogger().WithFields(mylog.String("bind_id", bindId)).
		Info("created a new bind id")
	s.id2data[bindId] = map[string]interface{}{
		"tel_a":   telA,
		"tel_b":   telX,
		"tel_x":   telB,
		"bind_id": bindId,
	}
	return &pb.BindReply{BindId: bindId}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (s *xbrServer) randomBindID() string {
	const bindIDLen = 10
	b := make([]rune, bindIDLen)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (s *xbrServer) Unbind(ctx context.Context, request *pb.UnbindRequest) (*pb.UnbindReply, error) {
	bindId := request.BindId
	mylog.CloneLogger().WithTag("app", "xbr-service").
		WithFields(mylog.String("bind_id", bindId)).
		Info("received unbind request")
	if _, ok := s.id2data[bindId]; ok {
		delete(s.id2data, bindId)
	} else {
		mylog.CloneLogger().Info("no such bind_id")
	}
	return &pb.UnbindReply{BindId: bindId}, nil
}
