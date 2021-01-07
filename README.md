# Go-Qibla

[![Go Report Card][report-badge]][report-url]
[![Go Reference][doc-badge]][doc-url]

Go-Qibla is a Go package for calculating qibla direction for a specific location.

## Usage Examples

Let's say we want to get qibla direction for Jakarta, Indonesia :

```go
package main

import (
	"fmt"

	"github.com/hablullah/go-qibla"
)

func main() {
	// Try to calculate qibla direction from Jakarta
	jktLatitude := -6.169777778
	jktLongitude := 106.8307333
	qiblaDirection, qiblaDistance := qibla.Get(jktLatitude, jktLongitude)
	fmt.Printf("DIRECTION: %f degrees\n", qiblaDirection)
	fmt.Printf("DISTANCE: %f km\n", qiblaDistance)
}
```

Which will give us following results :

```
DIRECTION: 295.143458 degrees
DISTANCE: 7918.900254 km
```

## Resources

1. Anugraha, R. 2012. _Mekanika Benda Langit_. ([PDF][pdf-rinto-anugraha])

## License

Go-Qibla is distributed using [MIT] license.

[report-badge]: https://goreportcard.com/badge/github.com/hablullah/go-qibla
[report-url]: https://goreportcard.com/report/github.com/hablullah/go-qibla
[doc-badge]: https://pkg.go.dev/badge/github.com/hablullah/go-qibla.svg
[doc-url]: https://pkg.go.dev/github.com/hablullah/go-qibla
[pdf-rinto-anugraha]: https://simpan.ugm.ac.id/s/GcxKuyZWn8Rshnn
[MIT]: http://choosealicense.com/licenses/mit/