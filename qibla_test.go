package qibla_test

import (
	"encoding/csv"
	"io"
	"math"
	"os"
	"strconv"
	"testing"

	"github.com/hablullah/go-qibla"
)

func Test_Qibla_Get(t *testing.T) {
	testData, err := openTestFile("test-data.csv")
	if err != nil {
		t.Fatal(err)
	}

	for _, data := range testData {
		qiblaDirection, qiblaDistance := qibla.Get(data.Latitude, data.Longitude)

		if math.Round((data.Direction-qiblaDirection)*1000) != 0 {
			t.Errorf("[%.0f, %.0f]: direction want %f got %f\n",
				data.Latitude,
				data.Longitude,
				data.Direction,
				qiblaDirection)
		}

		if math.Round((data.Distance-qiblaDistance)*1000) != 0 {
			t.Errorf("[%.0f, %.0f]: distance want %f got %f\n",
				data.Latitude,
				data.Longitude,
				data.Distance,
				qiblaDistance)
		}
	}
}

type TestData struct {
	Latitude  float64
	Longitude float64
	Direction float64
	Distance  float64
}

func openTestFile(csvFilePath string) ([]TestData, error) {
	// Open test file
	f, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Parse test file
	dataList := []TestData{}
	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		latitude, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			continue
		}

		longitude, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}

		direction, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			continue
		}

		distance, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}

		dataList = append(dataList, TestData{
			Latitude:  latitude,
			Longitude: longitude,
			Direction: direction,
			Distance:  distance,
		})
	}

	return dataList, nil
}
