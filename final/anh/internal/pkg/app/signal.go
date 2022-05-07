package app

import (
	"fmt"
	"os"
)

type SignalCancel struct {
	signal os.Signal
}

func (s *SignalCancel) Error() string {
	return fmt.Sprintf("cancelled by system signal %+v", s.signal)
}
