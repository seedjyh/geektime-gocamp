package goim

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_Pack(t *testing.T) {
	m := &Message{
		Header: Header{
			ProtocolVersion: 0x1112,
			Operation:       0x21222324,
			SequenceID:      0x31323334,
		},
		Body: []byte("Hello, world!"),
	}
	if buf, err := m.Pack(); assert.NoError(t, err) {
		assert.Equal(t, 0x1d, len(buf))
		assert.Equal(t, []byte{
			0x00, 0x00, 0x00, 0x1D, // package length (big endian)
			0x00, 0x10, // header length (big endian)
			0x11, 0x12, // protocol version (big endian)
			0x21, 0x22, 0x23, 0x24, // operation (big endian)
			0x31, 0x32, 0x33, 0x34, // sequence ID (big endian)
			// body: Hello, world!
			0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x77,
			0x6F, 0x72, 0x6C, 0x64, 0x21,
		}, buf)
	}
}
