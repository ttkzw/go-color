// Copyright 2021 Takashi Takizawa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/ttkzw/go-color"
)

const defaultColorName = "Reset"

var (
	colorName           string
	str                 string
	noTrailingLineBreak bool
)

func main() {
	flag.Parse()

	if colorName == "" {
		colorName = defaultColorName
	}

	if err := run(colorName, str, noTrailingLineBreak); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func init() {
	flag.BoolVar(&noTrailingLineBreak, "n", false, "Do not print the trailing newline character.")
	flag.StringVar(&colorName, "c", "", "color name.")
	flag.StringVar(&str, "s", "", "string")
}

func run(colorName string, str string, noTrailingLineBreak bool) error {
	c, err := color.NewColor(colorName)
	if err != nil {
		return err
	}

	var output = colorable.NewColorableStdout()

	s := c.Colorize(str)
	if noTrailingLineBreak {
		fmt.Fprint(output, s)
	} else {
		fmt.Fprintln(output, s)
	}
	return nil
}
