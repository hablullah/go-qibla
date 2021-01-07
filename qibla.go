package qibla

import (
	"math"
)

var (
	kaabaLatitude  = rad(21.42250833)
	kaabaLongitude = rad(39.82616111)
)

// Get returns the direction to qibla from specified coordinate.
func Get(latitude, longitude float64) float64 {
	// Convert coordinate to radian
	latitude = rad(latitude)
	longitude = rad(longitude)

	// Calculate direction
	y := math.Sin(kaabaLongitude - longitude)
	x := math.Cos(latitude)*math.Tan(kaabaLatitude) -
		math.Sin(latitude)*math.Cos(longitude-kaabaLongitude)

	direction := degree(math.Atan2(y, x))
	if direction < 0 {
		direction += 360
	}

	return direction
}

func rad(degree float64) float64 {
	return degree * math.Pi / 180
}

func degree(rad float64) float64 {
	return rad * 180 / math.Pi
}
