package main

import (
	"log"
	"os"

	"github.com/sauercrowd/http-perf/pkg/flags"
	"github.com/sauercrowd/http-perf/pkg/measure"
)

func main() {
	f := flags.Parse()
	var m measure.MeasurementResults
	// if amount > 0, do amount based measuring, otherwise time based
	if f.Amount > 0 {
		var err error
		m, err = measure.StartWithAmount(f.URL, f.Amount, f.GoroutineN)
		if err != nil {
			log.Println("Error: ", err)
			return
		}
	} else {
		var err error
		m, err = measure.StartWithTime(f.URL, f.TimeS, f.GoroutineN)
		if err != nil {
			log.Println("Error: ", err)
			return
		}
	}

	log.Printf("AVG request time: %f ms", m.GetAVG())
	if f.NoChart {
		return
	}

	//create graph
	bytes, err := m.GetChart()
	if err != nil {
		log.Println("Could create time chart: ", err)
		return
	}
	createGraph(bytes, f.ChartPath)
}

func createGraph(bytes []byte, path string) {
	chartF, err := os.Create(path)
	if err != nil {
		log.Println("Could open file: ", err)
		return
	}
	_, err = chartF.Write(bytes)
	if err != nil {
		log.Println("Could not write to file: ", err)
		return
	}
	if err := chartF.Close(); err != nil {
		log.Println("Could not close file: ", err)
		return
	}
	log.Println("Created Chart sucessfully at ", path)
}
