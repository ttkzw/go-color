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

### System Colors

- <span style="color:#000000;">SystemBlack</span>
- <span style="color:#AA0000;">SystemRed</span>
- <span style="color:#00AA00;">SystemGreen</span>
- <span style="color:#AAAA00;">SystemYellow</span>
- <span style="color:#0000AA;">SystemBlue</span>
- <span style="color:#AA00AA;">SystemMagenta</span>
- <span style="color:#00AAAA;">SystemCyan</span>
- <span style="color:#AAAAAA;">SystemWhite</span>
- <span style="color:#555555;">SystemBrightBlack</span>
- <span style="color:#FF5555;">SystemBrightRed</span>
- <span style="color:#55AA55;">SystemBrightGreen</span>
- <span style="color:#FFFF55;">SystemBrightYellow</span>
- <span style="color:#5555FF;">SystemBrightBlue</span>
- <span style="color:#FF55FF;">SystemBrightMagenta</span>
- <span style="color:#55FFFF;">SystemBrightCyan</span>
- <span style="color:#FFFFFF;">SystemBrightWhite</span>

### Colors

- <span style="color:#FF0000;">Red</span>
- <span style="color:#FF8700;">Orange</span>
- <span style="color:#FFFF00;">Yellow</span>
- <span style="color:#87FF00;">Chartreuse</span>
- <span style="color:#00FF00;">Green</span>
- <span style="color:#00FF87;">SpringGreen</span>
- <span style="color:#00FFFF;">Cyan</span>
- <span style="color:#0087FF;">Azure</span>
- <span style="color:#0000FF;">Blue</span>
- <span style="color:#8700FF;">Violet</span>
- <span style="color:#FF00FF;">Magenta</span>
- <span style="color:#FF0087;">Rose</span>
- <span style="color:#FFFFFF;">White</span>
- <span style="color:#808080;">Grey</span>
- <span style="color:#000000;">Black</span>

## Installation

```sh
$ go get github.com/ttkzw/go-color
```


## License

BSD 3-Clause License

## Auther

Takashi Takizawa
