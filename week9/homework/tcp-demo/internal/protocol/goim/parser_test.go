package goim

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	buf := []byte{
		0x00, 0x00, 0x00, 0x1D, // package length (big endian)
		0x00, 0x10, // header length (big endian)
		0x11, 0x12, // protocol version (big endian)
		0x21, 0x22, 0x23, 0x24, // operation (big endian)
		0x31, 0x32, 0x33, 0x34, // sequence ID (big endian)
		// body: Hello, world!
		0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x77,
		0x6F, 0x72, 0x6C, 0x64, 0x21,
		// extra bytes for next message...
		0xff, 0xff, 0xff, 0xff,
	}
	if m, l, err := NewParser().Parse(context.Background(), buf); assert.NoError(t, err) {
		assert.Equal(t, 0x1D, l)
		if m, ok := m.(*Message); assert.True(t, ok) {
			assert.Equal(t, int32(0x1112), m.Header.ProtocolVersion)
			assert.Equal(t, int32(0x21222324), m.Header.Operation)
			assert.Equal(t, int32(0x31323334), m.Header.SequenceID)
			assert.Equal(t, "Hello, world!", string(m.Body))
		}
	}
}
