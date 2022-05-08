package xbr

import (
	pb "anh/api"
	"anh/internal/pkg/mylog"
	"anh/internal/pkg/uuid"
	"context"
	"google.golang.org/grpc/metadata"
)

type xbrServer struct {
	pb.UnimplementedXBRServer
	id2data   map[string]interface{}
	bindIdGen uuid.Generator
}

func newXBRServer() *xbrServer {
	return &xbrServer{
		id2data:   make(map[string]interface{}),
		bindIdGen: uuid.NewUUID32Generator(),
	}
}

func (s *xbrServer) Bind(ctx context.Context, request *pb.BindRequest) (*pb.BindReply, error) {
	logger := mylog.CloneLogger().WithTag("app", "xbr-service")
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if v := md.Get("session_id"); len(v) > 0 {
			logger = logger.WithTag("session_id", v[0])
		}
	}
	telA := request.TelA
	telX := request.TelX
	telB := request.TelB
	logger.CloneLogger().
		WithFields(mylog.String("tel_a", telA)).
		WithFields(mylog.String("tel_x", telX)).
		WithFields(mylog.String("tel_b", telB)).
		Info("received bind request")
	bindId := s.bindIdGen.Next()
	logger.CloneLogger().WithFields(mylog.String("bind_id", bindId)).
		Info("created a new bind id")
	s.id2data[bindId] = map[string]interface{}{
		"tel_a":   telA,
		"tel_b":   telX,
		"tel_x":   telB,
		"bind_id": bindId,
	}
	return &pb.BindReply{BindId: bindId}, nil
}

func (s *xbrServer) Unbind(ctx context.Context, request *pb.UnbindRequest) (*pb.UnbindReply, error) {
	logger := mylog.CloneLogger().WithTag("app", "xbr-service")
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if v := md.Get("session_id"); len(v) > 0 {
			logger = logger.WithTag("session_id", v[0])
		}
	}
	bindId := request.BindId
	logger.CloneLogger().
		WithFields(mylog.String("bind_id", bindId)).
		Info("received unbind request")
	if _, ok := s.id2data[bindId]; ok {
		delete(s.id2data, bindId)
		logger.CloneLogger().
			WithFields(mylog.String("bind_id", bindId)).
			Info("found and unbind the binding")
	} else {
		logger.CloneLogger().
			WithFields(mylog.String("bind_id", bindId)).
			Info("no such bind_id")
	}
	return &pb.UnbindReply{BindId: bindId}, nil
}
