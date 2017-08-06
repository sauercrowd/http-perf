package measure

import (
	"context"
	"log"
	"sync"
	"time"
)

type TimeMeasurementResult struct {
	Value   int64
	Elapsed int64
}

type TimeMeasurement struct {
	Results []TimeMeasurementResult
}

func StartWithTime(url string, seconds int, goroutineN int, durationNS *int64) ([]MeasurementResult, error) {
	log.Printf("Measuring for %d seconds with %d goroutine(s)", seconds, goroutineN)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel() // release timing resources

	start := time.Now()
	var mtx sync.Mutex
	ch := make(chan MeasurementResult)
	for i := 0; i < goroutineN; i++ {
		go doTimeRequest(ctx, url, ch, start, &mtx)
	}

	var mr []MeasurementResult
	for v := range ch {
		mr = append(mr, v)
	}

	duration := time.Since(start).Nanoseconds()
	if durationNS != nil {
		*durationNS = duration
	}

	log.Printf("Took %f seconds for %d measurements", float64(duration)/float64(time.Second.Nanoseconds()), len(mr))
	return mr, nil
}

func doTimeRequest(ctx context.Context, url string, ch chan MeasurementResult, MeasurementStart time.Time, mtx *sync.Mutex) {
	for {
		d, err := doRequest(ctx, url, MeasurementStart)
		if err != nil {
			continue
		}
		//lock, so if ctx.Done() is propagated, channel will be closed and no one will try to use it anymore
		mtx.Lock()
		select {
		case <-ctx.Done():
			log.Println("Do not start another request")
			close(ch)
			return
		default:
			ch <- MeasurementResult{
				RequestTime: d.Elapsed.Nanoseconds(),
				SinceStart:  time.Since(MeasurementStart).Nanoseconds(),
				StatusCode:  d.StatusCode,
			}
		}
		mtx.Unlock()
	}
}
