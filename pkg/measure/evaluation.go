package measure

import (
	"bytes"
	"log"
	"time"

	"github.com/wcharczuk/go-chart"
)

func (tm *TimeMeasurement) GetAVG() float64 {
	var sum int64
	for _, v := range tm.Results {
		sum += v.Value
	}
	return float64(sum) / float64(len(tm.Results)) / float64(time.Millisecond)
}

func (tm *TimeMeasurement) GetChart() ([]byte, error) {
	//convert array to float array
	log.Printf("Got %d measurements", len(tm.Results))
	log.Println("Creating Chart...")
	resultsFloat64 := make([]float64, 0, len(tm.Results))
	msArr := make([]float64, 0, len(tm.Results))
	for _, r := range tm.Results {
		resultsFloat64 = append(resultsFloat64, float64(r.Value)/float64(time.Millisecond.Nanoseconds()))
		msArr = append(msArr, float64(r.Elapsed)/float64(time.Second.Nanoseconds()))
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true, //enables / displays the y-axis
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: msArr,
				YValues: resultsFloat64,
			},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	return buffer.Bytes(), err
}
