package qibla

import (
	"errors"

	"github.com/shopspring/decimal"
)

var (
	kaabaLatitude  = decimal.NewFromFloat(21.42250833)
	kaabaLongitude = decimal.NewFromFloat(39.82616111)

	// ErrIncalculable is error when qibla direction
	ErrIncalculabe = errors.New("qibla direction can't be calculated")
)

// Get returns the direction to qibla from specified location.
func Get(latitude float64, longitude float64) (float64, error) {
	// Convert latitude and longitude to decimal
	locationLatitude := decimal.NewFromFloat(latitude)
	locationLongitude := decimal.NewFromFloat(longitude)

	// Calculate qibla direction
	diffLongitude := kaabaLongitude.Sub(locationLongitude)
	B := cos(locationLatitude).Mul(tan(kaabaLatitude))
	C := sin(locationLatitude).Mul(cos(diffLongitude))
	Y := sin(diffLongitude)
	X := B.Sub(C)

	if X.Equal(decimal.Zero) {
		return 0, ErrIncalculabe
	}

	direction, _ := atan(Y.Div(X)).Float64()
	if direction < 0 {
		direction += 360.0
	}

	return direction, nil
}
