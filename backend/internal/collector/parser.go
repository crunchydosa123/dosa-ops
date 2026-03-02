package collector

import (
	"strconv"
	"strings"
)

func parseMetrics(data string) map[string]float64 {

	lines := strings.Split(data, "\n")

	metrics := make(map[string]float64)

	for _, line := range lines {

		parts := strings.Fields(line)

		if len(parts) != 2 {
			continue
		}

		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			continue
		}

		metrics[parts[0]] = value
	}

	return metrics
}
