package encoding

type bigEndian struct {
}

func (b *bigEndian) Int16(buf []byte) int16 {
	return int16(buf[0])<<8 | int16(buf[1])
}

func (b *bigEndian) Int32(buf []byte) int32 {
	return int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3])
}

func (b *bigEndian) PutInt16(buf []byte, v int16) {
	_ = buf[1] // 如果buf[1]越界，会发生panic。避免写越界
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
}

func (b *bigEndian) PutInt32(buf []byte, v int32) {
	_ = buf[3] // 如果buf[3]越界，会发生panic。避免写越界
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
}
