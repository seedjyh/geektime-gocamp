package xbr

import (
	pb "anh/api"
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
