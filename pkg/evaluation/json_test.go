package evaluation_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/sauercrowd/http-perf/pkg/evaluation"
)

const filepath = "test_json.json"

type testResultJSON struct {
	Requests       int         `json:"requests"`
	TimeTotal      int64       `json:"timeTotalNS"`
	TimePerRequest float64     `json:"timePerRequestMS"`
	ErrorCount     int         `json:"errorCount"`
	StatusCounts   map[int]int `json:"statusCounts"`
}

func TestJSON(t *testing.T) {
	statusMap := map[int]int{
		200: 14,
		500: 6,
	}
	if err := evaluation.WriteJSON(10, 20, 0.5, 2, statusMap, filepath); err != nil {
		t.Fatal("Could not write json:", err)
	}
	f, err := os.Open(filepath)
	if err != nil {
		t.Fatal("Could not open json file for checking:", err)
	}
	d := json.NewDecoder(f)
	var j testResultJSON
	if err := d.Decode(&j); err != nil {
		t.Fatal("Could not decode json for checking:", err)
	}
	//check values
	if j.Requests != 10 ||
		j.TimeTotal != 20 ||
		j.TimePerRequest != 0.5 {
		for k, v := range statusMap {
			value, ok := j.StatusCounts[k]
			if !ok {
				t.Error("Decoded json is not correct")
				t.Fatal("Got: ", j)
			}
			if value != v {
				t.Error("Decoded json is not correct")
				t.Fatal("Got: ", j)
			}
		}
		t.Error("Decoded json is not correct")
		t.Fatal("Got: ", j)
	}
	//remove json (not important for test, so log.Print used if it fails)
	if err := os.Remove(filepath); err != nil {
		t.Log("Could not remove ", filepath, err)
	}
}
