package channel

type Message interface {
	Pack() ([]byte, error)
}
