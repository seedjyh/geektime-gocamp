package channel

import (
	"context"
	"net"
)

type Sender struct {
	conn    net.Conn
	sending <-chan Message
}

func NewSender(conn net.Conn, sending <-chan Message) *Sender {
	return &Sender{
		conn:    conn,
		sending: sending,
	}
}

func (s *Sender) KeepSending(ctx context.Context) error {
	for m := range s.sending {
		if buf, err := m.Pack(); err != nil {
			return err
		} else if _, err := s.conn.Write(buf); err != nil {
			return err
		}
	}
	return nil
}
