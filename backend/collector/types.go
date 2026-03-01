package collector

import "time"

type Metric struct {
	ID        int
	Service   string
	Name      string
	Value     float64
	Timestamp time.Time
}
