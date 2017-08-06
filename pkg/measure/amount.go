package measure

import (
	"context"
	"log"
	"sync"
	"time"
)

func StartWithAmount(url string, count int, goroutineN int, durationNS *int64) ([]MeasurementResult, error) {
	log.Printf("Making %d requests with %d goroutine(s)", count, goroutineN)
	ch := make(chan MeasurementResult)
	ctx, cancel := context.WithCancel(context.Background())

	start := time.Now()
	var mtx sync.Mutex
	for i := 0; i < goroutineN; i++ {
		go doAmountRequest(ctx, url, ch, start, &mtx)
	}
	var mr []MeasurementResult
	i := 0
	for v := range ch {
		mr = append(mr, v)
		i++
		if i >= count {
			break
		}
	}
	cancel()
	duration := time.Since(start).Nanoseconds()
	//write time, if not nil
	if durationNS != nil {
		*durationNS = duration
	}
	log.Printf("Took %f seconds for %d measurements", float64(duration)/float64(time.Second.Nanoseconds()), count)
	return mr, nil
}

func doAmountRequest(ctx context.Context, url string, ch chan MeasurementResult, MeasurementStart time.Time, mtx *sync.Mutex) {
	for {
		d, err := doRequest(ctx, url, MeasurementStart)
		if err != nil {
			continue
		}
		//lock, so if ctx.Done() is propagated, channel will be closed and no one will try to use it anymore
		mtx.Lock()
		select {
		case <-ctx.Done():
			log.Println("Stopping")
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
