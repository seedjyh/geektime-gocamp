package rollingnumber

import "time"

type Clock interface {
	Now() time.Time
}
