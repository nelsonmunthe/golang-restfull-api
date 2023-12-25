//go:generate mockery --all
//go:generate mockery --all --inpackage --case snake

package clock

import (
	"time"
)

type clock struct {
}

func New() Clock {
	return clock{}
}

func (t clock) Now() time.Time {
	return time.Now()
}

func (t clock) Duration(val uint) uint {
	return uint(time.Duration(val))
}

func (t clock) Minutes() time.Duration {
	return time.Minute
}
