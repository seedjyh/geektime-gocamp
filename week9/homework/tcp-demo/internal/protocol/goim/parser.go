package goim

import (
	"context"
	"geektime-gocamp/week9/homework/tcp-demo/internal/pkg/channel"
	"geektime-gocamp/week9/homework/tcp-demo/internal/pkg/encoding"
	"github.com/pkg/errors"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

// Parse 解析 buf 并返回解析后的消息、解析消耗掉的字节数。
// 如果长度不足，返回 (nil, 0, *NeedMoreBytes)
func (p *Parser) Parse(ctx context.Context, buf []byte) (channel.Message, int, error) {
	if len(buf) < rawHeaderLength {
		return nil, 0, channel.NewNeedExtraBytes(rawHeaderLength - len(buf))
	}
	msg := new(Message)
	packLen := int(encoding.BigEndian.Int32(buf[packageLengthOffset:headerLengthOffset]))
	if len(buf) < packLen {
		return nil, 0, channel.NewNeedExtraBytes(packLen - len(buf))
	}
	headerLen := int(encoding.BigEndian.Int16(buf[headerLengthOffset:protocolVersionOffset]))
	if headerLen > packLen {
		return nil, 0, errors.New("Header length is larger than package length")
	}
	msg.Header.ProtocolVersion = int32(encoding.BigEndian.Int16(buf[protocolVersionOffset:operationOffset]))
	msg.Header.Operation = encoding.BigEndian.Int32(buf[operationOffset:sequenceOffset])
	msg.Header.SequenceID = encoding.BigEndian.Int32(buf[sequenceOffset:bodyOffset])
	if bodyLen := packLen - headerLen; bodyLen > 0 {
		msg.Body = make([]byte, bodyLen)
		copy(msg.Body, buf[bodyOffset:packLen])
	}
	return msg, packLen, nil
}
