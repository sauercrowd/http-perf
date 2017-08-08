package evaluation_test

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/sauercrowd/http-perf/pkg/measure"

	"github.com/sauercrowd/http-perf/pkg/evaluation"
)

func TestAVG(t *testing.T) {
	mr := []measure.MeasurementResult{
		measure.MeasurementResult{
			RequestTime: 5900,
			SinceStart:  55,
			StatusCode:  200,
		},
		measure.MeasurementResult{
			RequestTime: 1100,
			SinceStart:  55,
			StatusCode:  200,
		},
		measure.MeasurementResult{
			RequestTime: 1200,
			SinceStart:  55,
			StatusCode:  200,
		},
	}
	result := evaluation.GetAVG(mr)
	var shouldBe float64
	for _, m := range mr {
		shouldBe += float64(m.RequestTime)
	}
	shouldBe = shouldBe / float64(len(mr)) / float64(time.Millisecond)

	//checking
	if shouldBe != result {
		t.Error("Expected different result")
		t.Error("Expected: ", shouldBe)
		t.Fatal("Got:", result)
	}
}

func TestErrorCount(t *testing.T) {
	mr := []measure.MeasurementResult{
		measure.MeasurementResult{
			RequestTime: 5900,
			SinceStart:  55,
			StatusCode:  200,
		},
		measure.MeasurementResult{
			RequestTime: 1100,
			SinceStart:  55,
			StatusCode:  404,
		},
		measure.MeasurementResult{
			RequestTime: 1200,
			SinceStart:  55,
			StatusCode:  500,
		},
		measure.MeasurementResult{
			RequestTime: 100,
			SinceStart:  55,
			StatusCode:  200,
		},
	}
	result := evaluation.GetErrorCount(mr)
	var shouldBe int
	for _, m := range mr {
		if m.StatusCode != http.StatusOK {
			shouldBe++
		}
	}

	//checking
	if shouldBe != result {
		t.Error("Expected different result")
		t.Error("Expected: ", shouldBe)
		t.Fatal("Got:", result)
	}
}

func TestStatusCountMap(t *testing.T) {
	mr := []measure.MeasurementResult{
		measure.MeasurementResult{
			RequestTime: 5900,
			SinceStart:  55,
			StatusCode:  200,
		},
		measure.MeasurementResult{
			RequestTime: 1100,
			SinceStart:  55,
			StatusCode:  404,
		},
		measure.MeasurementResult{
			RequestTime: 1200,
			SinceStart:  55,
			StatusCode:  500,
		},
		measure.MeasurementResult{
			RequestTime: 100,
			SinceStart:  55,
			StatusCode:  200,
		},
	}
	expected := make(map[int]int)
	for _, k := range mr {
		if _, ok := expected[k.StatusCode]; ok {
			expected[k.StatusCode]++
			continue
		}
		expected[k.StatusCode] = 1
	}
	result := evaluation.GetStatusCountsMap(mr)
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		t.Error("Expected and result are not equal")
		t.Error("Expected:", expected)
		t.Fatal("Got:", result)
	}
}
