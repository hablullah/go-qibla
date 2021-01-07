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
	jktLatitude := -6.16978
	jktLongitude := 106.83073
	qiblaDirection := qibla.Get(jktLatitude, jktLongitude)
	fmt.Println(qiblaDirection)
}
```

Which will give us following results :

```
295.14345941710076
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