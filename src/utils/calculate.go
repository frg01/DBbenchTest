package utils

import (
	"sort"
	"time"
)

func CalculateP99(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	var totalDuration time.Duration
	for _, duration := range durations {
		totalDuration += duration
	}
	averageDuration := totalDuration / time.Duration(len(durations))
	_ = averageDuration

	var p99Index = int(float64(len(durations)) * 0.99)
	var sortedDurations = append([]time.Duration(nil), durations...)
	sort.Slice(sortedDurations, func(i, j int) bool { return sortedDurations[i] < sortedDurations[j] })

	return sortedDurations[p99Index]
}
