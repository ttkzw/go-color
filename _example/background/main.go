package main

import (
	"fmt"

	"github.com/ttkzw/go-color"
)

func main() {
	fmt.Println(color.GreenBackground.Colorize(color.White.Colorize("OK")))
}
