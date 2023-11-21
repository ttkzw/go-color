# go-color

This is a Go library for colorizing strings.

## Usage

```go
package main

import (
	"fmt"

	"github.com/ttkzw/go-color"
)

func main() {
    c, _ := color.NewColor("Green")
    fmt.Println(c.Colorize("OK"))
}
```

```go
package main

import (
	"fmt"

	"github.com/ttkzw/go-color"
)

func main() {
	fmt.Println(c.Green.Colorize("OK"))
}
```

If you want to support Windows, use go-colorable.

```go
import (
	"fmt"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/ttkzw/go-color"
)

func main() {
	var output = colorable.NewColorableStdout()

	fmt.Fprintln(output, c.Green.Colorize("OK"))
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
