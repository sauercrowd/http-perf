package main

import (
	"log"
	"os"

	"github.com/sauercrowd/http-perf/pkg/evaluation"
	"github.com/sauercrowd/http-perf/pkg/flags"
	"github.com/sauercrowd/http-perf/pkg/measure"
)

func main() {
	f := flags.Parse()
	var m []measure.MeasurementResult
	var durationNS int64
	// if amount > 0, do amount based measuring, otherwise time based
	if f.Amount > 0 {
		var err error
		m, err = measure.StartWithAmount(f.URL, f.Amount, f.GoroutineN, &durationNS)
		if err != nil {
			log.Println("Error:", err)
			return
		}
	} else {
		var err error
		m, err = measure.StartWithTime(f.URL, f.TimeS, f.GoroutineN, &durationNS)
		if err != nil {
			log.Println("Error:", err)
			return
		}
	}

	errorCount := evaluation.GetErrorCount(m)
	log.Printf("Got %d error(s)", errorCount)
	avg := evaluation.GetAVG(m)
	log.Printf("AVG request time: %f ms", avg)

	//create graph if needed
	if !f.NoChart {
		bytes, err := evaluation.GetChart(m)
		if err != nil {
			log.Println("Could create time chart:", err)
			return
		}
		createGraph(bytes, f.ChartPath)
	}

	//create JSON if needed
	if f.ResultJSON {
		statusCountMap := evaluation.GetStatusCountsMap(m)
		err := evaluation.WriteJSON(len(m), durationNS, avg, errorCount, statusCountMap, f.JSONPath)
		if err != nil {
			log.Println("Could not write json")
			return
		}
	}
}

func createGraph(bytes []byte, path string) {
	chartF, err := os.Create(path)
	if err != nil {
		log.Println("Could open file:", err)
		return
	}
	_, err = chartF.Write(bytes)
	if err != nil {
		log.Println("Could not write to file:", err)
		return
	}
	if err := chartF.Close(); err != nil {
		log.Println("Could not close file:", err)
		return
	}
	log.Println("Created Chart sucessfully at", path)
}
