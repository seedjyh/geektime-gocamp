package channel

import (
	"context"
	"fmt"
)

const MaxMessageLength = 1024

type Parser interface {
	// Parse 解析 buf 并返回解析后的消息、解析消耗掉的字节数。
	// 如果长度不足，返回 (nil, 0, *NeedMoreBytes)
	Parse(ctx context.Context, buf []byte) (Message, int, error)
}

type NeedExtraBytes struct {
	Length int
}

func NewNeedExtraBytes(Length int) *NeedExtraBytes {
	return &NeedExtraBytes{Length: Length}
}

func (n *NeedExtraBytes) Error() string {
	return fmt.Sprintf("need extra %d bytes", n.Length)
}
