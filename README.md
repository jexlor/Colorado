<img src="https://i.sstatic.net/sbSCk.png">


Simple and fast micro-library for colored text and symbols.

Colors supported: <strong>Black, Red, Green, Yellow, Blue, Purple, Cyan, White and more..</strong>

All of these colors can be accesed through `Color` function from <strong>colorado</strong> package

<h1>Example:</h1> 

```go
package main

import (
	"fmt"

	"github.com/jopiol/colorado"
)

func main() {
	text := "red text with bright blue background"
	fmt.Println(colorado.Color(text, colorado.Red, colorado.BrightBlueBg))
}

```
<h1>Example without background color:</h1> 

```go
func main() {
	text := "red text"
	fmt.Println(colorado.Color(text, colorado.Red, ""))
}
```

To install the package using `go get`, run the following command in your terminal:

``go get github.com/jopiol/colorado``
