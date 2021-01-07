package qibla

import (
	"math"
)

var (
	kaabaLatitude  = rad(21.42250833)
	kaabaLongitude = rad(39.82616111)
	equatorRadius  = 6378.137
	ellipsoidF     = 1.0 / 298.257222101
)

// Get returns the direction (in degrees) and distance (in km) to qibla
// from specified coordinate.
func Get(latitude, longitude float64) (float64, float64) {
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

	// Calculate distance
	u := ((kaabaLatitude + latitude) / 2)
	g := ((kaabaLatitude - latitude) / 2)
	j := ((kaabaLongitude - longitude) / 2)
	m := math.Pow(math.Sin(g), 2)*math.Pow(math.Cos(j), 2) +
		math.Pow(math.Cos(u), 2)*math.Pow(math.Sin(j), 2)
	n := math.Pow(math.Cos(g), 2)*math.Pow(math.Cos(j), 2) +
		math.Pow(math.Sin(u), 2)*math.Pow(math.Sin(j), 2)
	w := math.Atan(math.Sqrt(m / n))
	p := math.Sqrt(m*n) / w
	d := 2 * w * equatorRadius
	e1 := (3*p - 1) / (2 * n)
	e2 := (3*p + 1) / (2 * m)
	distance := d * (1 +
		ellipsoidF*e1*math.Pow(math.Sin(u), 2)*math.Pow(math.Cos(g), 2) -
		ellipsoidF*e2*math.Pow(math.Cos(u), 2)*math.Pow(math.Sin(g), 2))

	return direction, distance
}

func rad(degree float64) float64 {
	return degree * math.Pi / 180
}

func degree(rad float64) float64 {
	return rad * 180 / math.Pi
}
