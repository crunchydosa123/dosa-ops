package models

type MetricEvent struct {
	ServiceID int     `json:"service_id"`
	Service   string  `json:"service"`
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}
