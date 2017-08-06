package measure

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"
)

type MeasurementResult struct {
	RequestTime int64
	SinceStart  int64
	StatusCode  int
}

type requestResult struct {
	Elapsed    time.Duration
	StatusCode int
}

func doRequest(ctx context.Context, url string, MeasurementStart time.Time) (*requestResult, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	r = r.WithContext(ctx)
	r.Close = true
	if err != nil {
		log.Fatalln("Could create request:", err)
		return nil, err
	}
	start := time.Now()
	resp, err := http.DefaultClient.Do(r)
	elapsed := time.Since(start)
	if err != nil && strings.HasSuffix(err.Error(), ": context deadline exceeded") {
		return nil, nil
	}
	if err != nil {
		log.Fatalln("Could not do request:", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Got not okay status:", resp.StatusCode)
	}
	if err := resp.Body.Close(); err != nil {
		log.Println("Could not close body")
		return nil, err
	}
	return &requestResult{Elapsed: elapsed, StatusCode: resp.StatusCode}, nil
}
