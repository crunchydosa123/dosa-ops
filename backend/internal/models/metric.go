package models

import "time"

type Metric struct {
	Service   string
	Name      string
	Value     float64
	Timestamp time.Time
}
