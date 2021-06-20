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

<table>
<tbody>
<tr><td>Black</td><td>`#000000`</td></tr>
<tr><td>Red</td><td>`#AA0000`</td></tr>
<tr><td>Green</td><td>`#00AA00`</td></tr>
<tr><td>Yellow</td><td>`#AAAA00`</td></tr>
<tr><td>Blue</td><td>`#0000AA`</td></tr>
<tr><td>Magenta</td><td>`#AA00AA`</td></tr>
<tr><td>Cyan</td><td>`#00AAAA`</td></tr>
<tr><td>White</td><td>`#AAAAAA`</td></tr>
<tr><td>BrightBlack</td><td>`#555555`</td></tr>
<tr><td>BrightRed</td><td>`#FF5555`</td></tr>
<tr><td>BrightGreen</td><td>`#55AA55`</td></tr>
<tr><td>BrightYellow</td><td>`#FFFF55`</td></tr>
<tr><td>BrightBlue</td><td>`#5555FF`</td></tr>
<tr><td>BrightMagenta</td><td>`#FF55FF`</td></tr>
<tr><td>BrightCyan</td><td>`#55FFFF`</td></tr>
<tr><td>BrightWhite</td><td>`#FFFFFF`</td></tr>
</tbody>
</table>

### Colors

<table>
<tbody>
<tr><td>Red</td><td>`#FF0000`</td></tr>
<tr><td>Orange</td><td>`#FF8700`</td></tr>
<tr><td>Yellow</td><td>`#FFFF00`</td></tr>
<tr><td>Chartreuse</td><td>`#87FF00`</td></tr>
<tr><td>Green</td><td>`#00FF00`</td></tr>
<tr><td>SpringGreen</td><td>`#00FF87`</td></tr>
<tr><td>Cyan</td><td>`#00FFFF`</td></tr>
<tr><td>Azure</td><td>`#0087FF`</td></tr>
<tr><td>Blue</td><td>`#0000FF`</td></tr>
<tr><td>Violet</td><td>`#8700FF`</td></tr>
<tr><td>Magenta</td><td>`#FF00FF`</td></tr>
<tr><td>Rose</td><td>`#FF0087`</td></tr>
<tr><td>White</td><td>`#FFFFFF`</td></tr>
<tr><td>Grey</td><td>`#808080`</td></tr>
<tr><td>Black</td><td>`#000000`</td></tr>
</tbody>
</table>

## Installation

```sh
$ go get github.com/ttkzw/go-color
```


## License

BSD 3-Clause License

## Auther

Takashi Takizawa
