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
