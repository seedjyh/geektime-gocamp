package rollingnumber

import "time"

type SystemClock struct{}

func (s SystemClock) Now() time.Time {
	return time.Now()
}
