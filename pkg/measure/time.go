package measure

import (
	"context"
	"fmt"
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

func StartWithTime(url string, seconds int, goroutineN int) ([]MeasurementResult, error) {
	log.Printf("Measuring for %d seconds with %d goroutines", seconds, goroutineN)
	timeout, err := time.ParseDuration(fmt.Sprintf("%ds", seconds))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())

	start := time.Now()
	var mtx sync.Mutex
	ch := make(chan MeasurementResult)

	for i := 0; i < goroutineN; i++ {
		go doTimeRequest(ctx, url, ch, start, &mtx)
	}
	//cancel after x seconds
	go func() {
		time.Sleep(timeout)
		cancel()
	}()
	var mr []MeasurementResult
	for v := range ch {
		mr = append(mr, v)
	}
	log.Printf("Took %f seconds for %d measurements", float64(time.Since(start).Nanoseconds())/float64(time.Second.Nanoseconds()), len(mr))
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
			ch <- MeasurementResult{Value: d.Nanoseconds(), Elapsed: time.Since(MeasurementStart).Nanoseconds()}
		}
		mtx.Unlock()
	}
}
