package flags

import (
	"flag"
)

type Flags struct {
	TimeS, Amount, GoroutineN int
	URL, ChartPath            string
	NoChart                   bool
}

func Parse() Flags {
	var f Flags
	flag.StringVar(&f.URL, "url", "", "url which should be measured")
	flag.IntVar(&f.TimeS, "time", 60, "how many seconds should be measured. Will be ignored if amount flag is set")
	flag.IntVar(&f.Amount, "amount", 0, "how many times shoud be measured. If used, time flag will be ignored")
	flag.IntVar(&f.GoroutineN, "n", 100, "amount of goroutines beeing used")
	flag.BoolVar(&f.NoChart, "nochart", false, "determines if a chart should be generated")
	flag.StringVar(&f.ChartPath, "chartpath", "perf.png", "path for chart png")
	flag.Parse()
	return f
}
