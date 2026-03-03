package models

import "time"

type MetricPoint struct {
	Timestamp time.Time
	Value     float64
}
