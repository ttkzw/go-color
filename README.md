# go-color

This is a Go library for colorizing strings.

## Usage

```go
package main

import (
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
	"github.com/ttkzw/go-color"
)

func main() {
    fmt.Println(c.Green.Colorize("OK"))
}
```

## Supported colors

### System Colors

- Black `#000000`
- Red `#AA0000`
- Green `#00AA00`
- Yellow `#AAAA00`
- Blue `#0000AA`
- Magenta `#AA00AA`
- Cyan `#00AAAA`
- White `#AAAAAA`
- BrightBlack `#555555`
- BrightRed `#FF5555`
- BrightGreen `#55AA55`
- BrightYellow `#FFFF55`
- BrightBlue `#5555FF`
- BrightMagenta `#FF55FF`
- BrightCyan `#55FFFF`
- BrightWhite `#FFFFFF`

### Colors

- Red `#FF0000`
- Orange `#FF8700`
- Yellow `#FFFF00`
- Chartreuse `#87FF00`
- Green `#00FF00`
- SpringGreen `#00FF87`
- Cyan `#00FFFF`
- Azure `#0087FF`
- Blue `#0000FF`
- Violet `#8700FF`
- Magenta `#FF00FF`
- Rose `#FF0087`
- White `#FFFFFF`
- Grey `#808080`
- Black `#000000`

## Installation

```sh
$ go get github.com/ttkzw/go-color
```


## License

BSD 3-Clause License

## Auther

Takashi Takizawa
