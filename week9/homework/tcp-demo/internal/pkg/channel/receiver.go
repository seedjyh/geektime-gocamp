package channel

import (
	"context"
	"github.com/pkg/errors"
	"net"
)

type Receiver struct {
	conn     net.Conn
	parser   Parser
	received chan Message
}

func NewReceiver(conn net.Conn, parser Parser) *Receiver {
	return &Receiver{
		conn:     conn,
		parser:   parser,
		received: make(chan Message),
	}
}

func (r Receiver) KeepReceiving(ctx context.Context) error {
	defer close(r.received)
	buf := make([]byte, MaxMessageLength)
	nowLength := 0
	for {
		if n, err := r.conn.Read(buf[nowLength:]); err != nil {
			return err
		} else {
			nowLength += n
		}
		if message, length, err := r.parser.Parse(ctx, buf[:nowLength]); err != nil {
			e := new(NeedExtraBytes)
			if errors.As(err, &e) {
				continue
			} else {
				return err
			}
		} else {
			r.received <- message
			newBuf := make([]byte, MaxMessageLength)
			copy(newBuf, buf[length:])
			buf = newBuf
			nowLength -= length
		}
	}
}

func (r *Receiver) ReceivedMessage() <-chan Message {
	return r.received
}
