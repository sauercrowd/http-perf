package measure

import (
	"bytes"
	"log"
	"time"

	"github.com/wcharczuk/go-chart"
)

type MeasurementResults []MeasurementResult

func (mr MeasurementResults) GetAVG() float64 {
	var sum int64
	for _, v := range mr {
		sum += v.Value
	}
	return float64(sum) / float64(len(mr)) / float64(time.Millisecond)
}

func (mr MeasurementResults) GetChart() ([]byte, error) {
	//convert array to float array
	log.Println("Creating Chart...")
	resultsFloat64 := make([]float64, 0, len(mr))
	msArr := make([]float64, 0, len(mr))
	for _, r := range mr {
		resultsFloat64 = append(resultsFloat64, float64(r.Value)/float64(time.Millisecond.Nanoseconds()))
		msArr = append(msArr, float64(r.Elapsed)/float64(time.Second.Nanoseconds()))
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Time elapsed (s)",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
		},
		YAxis: chart.YAxis{
			Name:      "Time server response (ms)",
			NameStyle: chart.StyleShow(),
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
