package evaluation

import (
	"net/http"
	"time"

	"github.com/sauercrowd/http-perf/pkg/measure"
)

// GetAVG returns the average time spent on the requests
func GetAVG(mr []measure.MeasurementResult) float64 {
	var sum int64
	for _, r := range mr {
		sum += r.RequestTime
	}
	// sum divided by amount divided by time.Milliseconds to get milliseconds
	return float64(sum) / float64(len(mr)) / float64(time.Millisecond)
}

// GetErrorCount returns the count of non-200 http status codes
func GetErrorCount(mr []measure.MeasurementResult) int {
	var count int
	for _, r := range mr {
		if r.StatusCode != http.StatusOK {
			count++
		}
	}
	return count
}

// GetStatusCountsMap returns a map with http statusCodes as keys and a count, how often this code appeared, as value
func GetStatusCountsMap(mr []measure.MeasurementResult) map[int]int {
	statusCodeCounts := make(map[int]int)
	for _, r := range mr {
		statusCodeCounts[r.StatusCode]++
	}
	return statusCodeCounts
}
