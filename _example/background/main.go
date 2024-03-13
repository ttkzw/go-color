package main

import (
	"fmt"

	"github.com/ttkzw/go-color"
)

func main() {
	fmt.Println(color.Colorize("OK", color.White, color.GreenBackground))
}
