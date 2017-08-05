package measure

import (
	"context"
	"log"
	"net/http"
	"time"
)

type MeasurementResult struct {
	Value   int64
	Elapsed int64
}

func doRequest(ctx context.Context, url string, MeasurementStart time.Time) (*time.Duration, error) {
	var client http.Client
	r, err := http.NewRequest("GET", url, nil)
	r.WithContext(ctx)
	r.Close = true
	if err != nil {
		log.Fatalln("Could create request: ", err)
		return nil, err
	}
	start := time.Now()
	resp, err := client.Do(r)
	elapsed := time.Since(start)
	if err != nil {
		log.Fatalln("Could not do request: ", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Got not okay status: ", resp.StatusCode)
	}
	if err := resp.Body.Close(); err != nil {
		log.Println("Could not close body")
		return nil, err
	}
	return &elapsed, nil
}
