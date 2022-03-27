package channel

import (
	"context"
	"net"
	"sync"
)

type Channel struct {
	conn     net.Conn
	received chan Message
	sending  chan Message
	receiver *Receiver
}

func NewChannel(conn net.Conn, parser Parser) *Channel {
	return &Channel{
		conn:     conn,
		received: make(chan Message),
		sending:  make(chan Message),
		receiver: NewReceiver(conn, parser),
	}
}

func (ch *Channel) Start(ctx context.Context) {
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = ch.receiver.KeepReceiving(ctx)
	}()
	go func() {
		defer wg.Done()
		_ = NewSender(ch.conn, ch.sending).KeepSending(ctx)
	}()
}

func (ch *Channel) Stop() {
	_ = ch.conn.Close()

}

func (ch *Channel) Receive() <-chan Message {
	return ch.receiver.ReceivedMessage()
}
