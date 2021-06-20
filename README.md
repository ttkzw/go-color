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
<tr><td>Black #000000</td></tr>
<tr><td>Red #AA0000</td></tr>
<tr><td>Green #00AA00</td></tr>
<tr><td>Yellow #AAAA00</td></tr>
<tr><td>Blue #0000AA</td></tr>
<tr><td>Magenta #AA00AA</td></tr>
<tr><td>Cyan #00AAAA</td></tr>
<tr><td>White #AAAAAA</td></tr>
<tr><td>BrightBlack #555555</td></tr>
<tr><td>BrightRed #FF5555</td></tr>
<tr><td>BrightGreen #55AA55</td></tr>
<tr><td>BrightYellow #FFFF55</td></tr>
<tr><td>BrightBlue #5555FF</td></tr>
<tr><td>BrightMagenta #FF55FF</td></tr>
<tr><td>BrightCyan #55FFFF</td></tr>
<tr><td>BrightWhite #FFFFFF</td></tr>
</tbody>
</table>

### Colors

<table>
<tbody>
<tr><td>Red #FF0000</td></tr>
<tr><td>Orange #FF8700</td></tr>
<tr><td>Yellow #FFFF00</td></tr>
<tr><td>Chartreuse #87FF00</td></tr>
<tr><td>Green #00FF00</td></tr>
<tr><td>SpringGreen #00FF87</td></tr>
<tr><td>Cyan #00FFFF</td></tr>
<tr><td>Azure #0087FF</td></tr>
<tr><td>Blue #0000FF</td></tr>
<tr><td>Violet #8700FF</td></tr>
<tr><td>Magenta #FF00FF</td></tr>
<tr><td>Rose #FF0087</td></tr>
<tr><td>White #FFFFFF</td></tr>
<tr><td>Grey #808080</td></tr>
<tr><td>Black #000000</td></tr>
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
