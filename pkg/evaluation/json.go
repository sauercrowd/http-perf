package evaluation

import (
	"encoding/json"
	"log"
	"os"
)

type resultJSON struct {
	Requests       int         `json:"requests"`
	TimeTotal      int64       `json:"timeTotalNS"`
	TimePerRequest float64     `json:"timePerRequestMS"`
	ErrorCount     int         `json:"errorCount"`
	StatusCounts   map[int]int `json:"statusCounts"`
}

// WriteJSON takes all parameters and write it to a json at filepath
func WriteJSON(requests int, timeTotal int64, timePerRequest float64, errorCount int, statusCounts map[int]int, filepath string) error {
	r := resultJSON{
		Requests:       requests,
		TimeTotal:      timeTotal,
		TimePerRequest: timePerRequest,
		ErrorCount:     errorCount,
		StatusCounts:   statusCounts,
	}
	bytes, err := json.Marshal(&r)
	if err != nil {
		log.Println("Could not marshal result struct: ", err)
		return err
	}
	f, err := os.Create(filepath)
	if err != nil {
		log.Println("Could not open file: ", err)
		return err
	}
	_, err = f.Write(bytes)
	if err != nil {
		log.Println("Could not write to file: ", err)
		return err
	}
	return nil
}
