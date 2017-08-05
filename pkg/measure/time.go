package measure

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

func StartWithTime(url string, seconds int, goroutineN int) (*TimeMeasurement, error) {
	ch := make(chan TimeMeasurementResult)
	log.Printf("Measuring for %d seconds with %d goroutines", seconds, goroutineN)
	timeout, err := time.ParseDuration(fmt.Sprintf("%ds", seconds))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	start := time.Now()
	doRequests(ctx, url, ch, start, goroutineN)
	go func() {
		time.Sleep(timeout)
		cancel()
	}()
	var tm TimeMeasurement
	for v := range ch {
		tm.Results = append(tm.Results, v)
	}
	return &tm, nil
}

func doRequests(ctx context.Context, url string, ch chan TimeMeasurementResult, MeasurementStart time.Time, goroutineN int) {
	var mtx sync.Mutex
	for i := 0; i < goroutineN; i++ {
		select {
		case <-ctx.Done():
			log.Println("Finishing")
			return
		default:
		}
		go func() {
			for {
				doRequest(ctx, url, ch, &mtx, MeasurementStart)
			}
		}()
	}
}

func doRequest(ctx context.Context, url string, ch chan TimeMeasurementResult, mtx *sync.Mutex, MeasurementStart time.Time) {
	var client http.Client
	r, err := http.NewRequest("GET", url, nil)
	r.WithContext(ctx)
	r.Close = true
	if err != nil {
		log.Fatalln("Could create request: ", err)
	}
	start := time.Now()
	resp, err := client.Do(r)
	elapsed := time.Since(start)
	if err != nil {
		log.Fatalln("Could not do request: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Got not okay status: ", resp.StatusCode)
	}
	if err := resp.Body.Close(); err != nil {
		log.Println("Could not close body")
		return
	}
	//lock, so if ctx.Done() is propagated, channel will be closed and no one will try to use it anymore
	mtx.Lock()
	select {
	case <-ctx.Done():
		log.Println("Do not start another request")
		close(ch)
		return
	default:
		e := time.Since(MeasurementStart)
		ch <- TimeMeasurementResult{Value: elapsed.Nanoseconds(), Elapsed: e.Nanoseconds()}
	}
	mtx.Unlock()
}
