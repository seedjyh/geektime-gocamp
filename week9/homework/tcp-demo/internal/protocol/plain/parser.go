package plain

import (
	"context"
	"geektime-gocamp/week9/homework/tcp-demo/internal/pkg/channel"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(ctx context.Context, buf []byte) (channel.Message, int, error) {
	return &Message{b: buf}, len(buf), nil
}
