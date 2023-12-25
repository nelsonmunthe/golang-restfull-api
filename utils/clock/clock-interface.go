package clock

import (
	"time"
)

type Clock interface {
	Now() time.Time
	Duration(val uint) uint
	Minutes() time.Duration
}
