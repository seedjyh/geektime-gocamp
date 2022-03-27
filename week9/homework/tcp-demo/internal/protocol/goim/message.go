package goim

import "geektime-gocamp/week9/homework/tcp-demo/internal/pkg/encoding"

const (
	packageLengthOffset   = 0
	packageLengthSize     = 4
	headerLengthOffset    = packageLengthOffset + packageLengthSize
	headerLengthSize      = 2
	protocolVersionOffset = headerLengthOffset + headerLengthSize
	protocolVersionSize   = 2
	operationOffset       = protocolVersionOffset + protocolVersionSize
	operationSize         = 4
	sequenceOffset        = operationOffset + operationSize
	sequenceSize          = 4
	bodyOffset            = sequenceOffset + sequenceSize

	rawHeaderLength = bodyOffset
)

type Header struct {
	ProtocolVersion int32
	Operation       int32
	SequenceID      int32
}

type Message struct {
	Header Header
	Body   []byte
}

func (m *Message) Pack() ([]byte, error) {
	headerLen := rawHeaderLength
	packLen := headerLen + len(m.Body)
	buf := make([]byte, packLen)
	encoding.BigEndian.PutInt32(buf[packageLengthOffset:], int32(packLen))
	encoding.BigEndian.PutInt16(buf[headerLengthOffset:], int16(headerLen))
	encoding.BigEndian.PutInt16(buf[protocolVersionOffset:], int16(m.Header.ProtocolVersion))
	encoding.BigEndian.PutInt32(buf[operationOffset:], int32(m.Header.Operation))
	encoding.BigEndian.PutInt32(buf[sequenceOffset:], int32(m.Header.SequenceID))
	copy(buf[bodyOffset:], m.Body)
	return buf, nil
}
