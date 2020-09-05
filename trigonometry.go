package qibla

import (
	"math"

	"github.com/shopspring/decimal"
)

var (
	decPi  = decimal.NewFromFloat(math.Pi)
	dec180 = decimal.New(180, 0)
)

// sin calculate sin value from the specified degree.
func sin(d decimal.Decimal) decimal.Decimal {
	rad, _ := d.Mul(decPi).Div(dec180).Float64()
	return decimal.NewFromFloat(math.Sin(rad))
}

// cos calculate cos value from the specified degree.
func cos(d decimal.Decimal) decimal.Decimal {
	rad, _ := d.Mul(decPi).Div(dec180).Float64()
	return decimal.NewFromFloat(math.Cos(rad))
}

// tan calculate tan value from the specified degree.
func tan(d decimal.Decimal) decimal.Decimal {
	rad, _ := d.Mul(decPi).Div(dec180).Float64()
	return decimal.NewFromFloat(math.Tan(rad))
}

// atan calculate degree from specified tan value.
func atan(tanValue decimal.Decimal) decimal.Decimal {
	fl, _ := tanValue.Float64()

	return decimal.NewFromFloat(math.Atan(fl)).
		Mul(dec180).
		Div(decPi)
}
