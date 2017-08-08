package measure_test

import (
	"log"
	"testing"
	"time"

	"github.com/sauercrowd/http-perf/pkg/measure"
)

func TestStartWithAmount(t *testing.T) {
	var duration int64
	res, err := measure.StartWithAmount("http://github.com", 10, 1, &duration)
	if err != nil {
		log.Fatal("Error happened while testing:", err)
	}
	t.Logf("Took %f miliseconds for 10 requests", float64(duration)/float64(time.Millisecond))
	t.Logf("First Request:\n\t RequestTime: %d\n\tSinceStart: %d\n\tStatusCode: %d", res[0].RequestTime, res[0].SinceStart, res[0].StatusCode)
}
