# go-color

This is a Go library for colorizing strings.

## Usage

To give a color name, write as follows:

```go
package main

import (
	"fmt"
	"log"

	"github.com/ttkzw/go-color"
)

func main() {
	c, err := color.NewColor("Green")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Colorize("OK"))
}
```

To use built-in constants, write as follows:

```go
package main

import (
	"fmt"

	"github.com/ttkzw/go-color"
)

func main() {
	fmt.Println(color.Green.Colorize("OK"))
}
```

To use a background color, write as follows:

```go
package main

import (
	"fmt"

	"github.com/ttkzw/go-color"
)

func main() {
	fmt.Println(color.GreenBackground.Colorize(color.White.Colorize("OK")))
}
```

If you want to support Windows, use go-colorable.

```go
package main

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/ttkzw/go-color"
)

func main() {
	var output = colorable.NewColorableStdout()

	fmt.Fprintln(output, color.Green.Colorize("OK"))
}
```


## Supported colors

- [Supported colors](colors.md)

## Installation

```sh
$ go get github.com/ttkzw/go-color
```


## License

BSD 3-Clause License

## Auther

Takashi Takizawa
