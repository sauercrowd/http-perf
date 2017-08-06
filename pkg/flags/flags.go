package flags

import (
	"flag"
	"log"
	"net/url"
)

// Flags is used for command line flags
type Flags struct {
	TimeS, Amount, GoroutineN int
	URL, ChartPath, JSONPath  string
	NoChart, ResultJSON       bool
}

// Parse parses the command line flags and returns a Flags struct
func Parse() Flags {
	var f Flags
	flag.StringVar(&f.URL, "url", "", "url which should be measured")
	flag.IntVar(&f.TimeS, "time", 60, "how many seconds should be measured. Will be ignored if amount flag is set")
	flag.IntVar(&f.Amount, "count", 0, "how many times shoud be measured. If used, time flag will be ignored")
	flag.IntVar(&f.GoroutineN, "n", 100, "amount of goroutines beeing used")
	flag.BoolVar(&f.NoChart, "nochart", false, "set if no chart should be generated")
	flag.StringVar(&f.ChartPath, "chartpath", "perf.png", "path for chart png")
	flag.BoolVar(&f.ResultJSON, "json", false, "creates a json file with keys requests, timeTotalNS, timePerRequestMS, errorCount(non 200 http status), statusCounts(object, httpstatuscode as key, count as value)")
	flag.StringVar(&f.JSONPath, "jsonpath", "perf.json", "filepath for json file")
	flag.Parse()
	checkURL(f.URL)
	return f
}

func checkURL(u string) {
	if u == "" {
		log.Fatal("Please pass a URL")
	}
	_, err := url.ParseRequestURI(u)
	if err != nil {
		log.Fatalf("URL %s is not valid", u)
	}
}
