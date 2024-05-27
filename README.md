Simple and fast micro-library for colored text and symbols.

Colors supported: <strong>Black, Red, Green, Yellow, Blue, Purple, Cyan, White.</strong>

All of these colors can be accesed through `Colorado` function

Example with code: 

```go
package main

import (
	"fmt"

	"github.com/GeorgeStoic/colorado"
)

func main() {
	text := "red text"
	fmt.Println(colorado.Colorado(text, colorado.Red))
}

```
To install the package using `go get`, run the following command in your terminal:

``go get github.com/GeorgeStoic/colorado``
