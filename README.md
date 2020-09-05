# Go-Qibla

[![Go Report Card](https://goreportcard.com/badge/github.com/RadhiFadlillah/go-qibla)](https://goreportcard.com/report/github.com/RadhiFadlillah/go-qibla)
[![GoDoc](https://godoc.org/github.com/RadhiFadlillah/go-qibla?status.png)](https://godoc.org/github.com/RadhiFadlillah/go-qibla)

Go-Qibla is a Go package for calculating qibla direction for a specific location.

## Usage Examples

For example, we want to get qibla direction for Jakarta, Indonesia :

```go
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
```

Which will give us following results :

```
295.14345941710076
```

## Resources

1. Anugraha, R. 2012. _Mekanika Benda Langit_. ([PDF](https://simpan.ugm.ac.id/s/GcxKuyZWn8Rshnn))

## License

Go-Qibla is distributed using [MIT](http://choosealicense.com/licenses/mit/) license.
