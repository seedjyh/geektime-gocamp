package plain

type Message struct {
	b []byte
}

func (m *Message) Pack() ([]byte, error) {
	return m.b, nil
}
