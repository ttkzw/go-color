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
<tr><td>Black</td><td><span style="color:#000000;">■</td></tr>
<tr><td>Red</td><td><span style="color:#AA0000;">■</td></tr>
<tr><td>Green</td><td><span style="color:#00AA00;">■</td></tr>
<tr><td>Yellow</td><td><span style="color:#AAAA00;">■</td></tr>
<tr><td>Blue</td><td><span style="color:#0000AA;">■</td></tr>
<tr><td>Magenta</td><td><span style="color:#AA00AA;">■</td></tr>
<tr><td>Cyan</td><td><span style="color:#00AAAA;">■</td></tr>
<tr><td>White</td><td><span style="color:#AAAAAA;">■</td></tr>
<tr><td>BrightBlack</td><td><span style="color:#555555;">■</td></tr>
<tr><td>BrightRed</td><td><span style="color:#FF5555;">■</td></tr>
<tr><td>BrightGreen</td><td><span style="color:#55AA55;">■</td></tr>
<tr><td>BrightYellow</td><td><span style="color:#FFFF55;">■</td></tr>
<tr><td>BrightBlue</td><td><span style="color:#5555FF;">■</td></tr>
<tr><td>BrightMagenta</td><td><span style="color:#FF55FF;">■</td></tr>
<tr><td>BrightCyan</td><td><span style="color:#55FFFF;">■</td></tr>
<tr><td>BrightWhite</td><td><span style="color:#FFFFFF;">■</td></tr>
</tbody>
</table>

### Colors

<table>
<tbody>
<tr><td>Red</td><td><span style="color:#FF0000;">■</td></tr>
<tr><td>Orange</td><td><span style="color:#FF8700;">■</td></tr>
<tr><td>Yellow</td><td><span style="color:#FFFF00;">■</td></tr>
<tr><td>Chartreuse</td><td><span style="color:#87FF00;">■</td></tr>
<tr><td>Green</td><td><span style="color:#00FF00;">■</td></tr>
<tr><td>SpringGreen</td><td><span style="color:#00FF87;">■</td></tr>
<tr><td>Cyan</td><td><span style="color:#00FFFF;">■</td></tr>
<tr><td>Azure</td><td><span style="color:#0087FF;">■</td></tr>
<tr><td>Blue</td><td><span style="color:#0000FF;">■</td></tr>
<tr><td>Violet</td><td><span style="color:#8700FF;">■</td></tr>
<tr><td>Magenta</td><td><span style="color:#FF00FF;">■</td></tr>
<tr><td>Rose</td><td><span style="color:#FF0087;">■</td></tr>
<tr><td>White</td><td><span style="color:#FFFFFF;">■</td></tr>
<tr><td>Grey</td><td><span style="color:#808080;">■</td></tr>
<tr><td>Black</td><td><span style="color:#000000;">■</td></tr>
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
