package component

import (
	"time"
)

type Tick struct {
	Interval  time.Duration
	Creation  time.Time
	EOL       time.Time
	TickEvent func()
	EOLEvent  func()
}
