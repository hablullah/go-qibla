// +build ignore

package main

import (
	"fmt"

	"github.com/RadhiFadlillah/go-qibla"
)

func main() {
	// Try to calculate qibla direction from Jakarta
	jktLatitude := -6.16978
	jktLongitude := 106.83073
	qiblaDirection, _ := qibla.Get(jktLatitude, jktLongitude)
	fmt.Println(qiblaDirection)
}
