package encoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigEndian_Int16(t *testing.T) {
	assert.Equal(t, int16(0x1234), BigEndian.Int16([]byte{0x12, 0x34}))
}

func TestBigEndian_Int32(t *testing.T) {
	assert.Equal(t, int32(0x12345678), BigEndian.Int32([]byte{0x12, 0x34, 0x56, 0x78}))
}
