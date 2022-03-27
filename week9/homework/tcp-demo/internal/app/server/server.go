package server

import (
	"context"
	"errors"
	"fmt"
	"geektime-gocamp/week9/homework/tcp-demo/internal/pkg/channel"
	"geektime-gocamp/week9/homework/tcp-demo/internal/protocol/goim"
	"net"
	"sync"
)

type Server struct {
	addr     string
	listener net.Listener
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Start(ctx context.Context) error {
	if ln, err := net.Listen("tcp", s.addr); err != nil {
		return err
	} else {
		s.listener = ln
	}
	var wg sync.WaitGroup
	defer wg.Wait()
	serverCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		if conn, err := s.listener.Accept(); err != nil {
			return err
		} else {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_ = s.handleConnection(serverCtx, conn)
			}()
		}
	}
}

func (s *Server) Stop(ctx context.Context) error {
	return s.listener.Close()
}

// handleConnection 阻塞地处理一个连接。当 conn 出错，或者ctx被cancel，则关闭连接并返回。
func (s *Server) handleConnection(ctx context.Context, conn net.Conn) error {
	var wg sync.WaitGroup
	defer wg.Wait()
	defer func() {
		_ = conn.Close()
	}()
	// receiver
	receiver := channel.NewReceiver(conn, goim.NewParser())
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = receiver.KeepReceiving(ctx)
	}()
	// receiver
	sending := make(chan channel.Message)
	// defer close(sending) 为防万一，不关闭这个channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = channel.NewSender(conn, sending).KeepSending(ctx)
	}()
	// processing
	for {
		select {
		case <-ctx.Done():
			_ = conn.Close()
			return errors.New("context cancelled")
		case msg := <-receiver.ReceivedMessage():
			fmt.Println("Received message", msg)
			fmt.Println("Sending message", msg)
			sending <- msg
		}
	}
}
