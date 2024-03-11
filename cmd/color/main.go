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

var (
	colorName           string
	backgroundColorName string
	str                 string
	noTrailingLineBreak bool
	listColors          bool
)

func main() {
	flag.Parse()

	if listColors {
		list()
	} else {
		if err := run(colorName, backgroundColorName, str, noTrailingLineBreak); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
	os.Exit(0)
}

func init() {
	flag.BoolVar(&noTrailingLineBreak, "n", false, "Do not print the trailing newline character.")
	flag.StringVar(&colorName, "c", "", "color name.")
	flag.StringVar(&backgroundColorName, "b", "", "background color name.")
	flag.StringVar(&str, "s", "", "string")
	flag.BoolVar(&listColors, "l", false, "Displays a list of supported colors.")
}

func run(colorName, backgroundColorName, str string, noTrailingLineBreak bool) error {
	var s string

	if colorName != "" {
		c, err := color.NewColor(colorName)
		if err != nil {
			return err
		}
		s = c.Colorize(str)
	} else {
		s = str
	}

	if backgroundColorName != "" {
		b, err := color.NewBackgroundColor(backgroundColorName)
		if err != nil {
			return err
		}
		s = b.Colorize(s)
	}

	var output = colorable.NewColorableStdout()
	if noTrailingLineBreak {
		fmt.Fprint(output, s)
	} else {
		fmt.Fprintln(output, s)
	}
	return nil
}

func list() {
	colors := color.SupportedColors()
	for _, c := range colors {
		s := c.Colorize(c.String())
		fmt.Printf("%s\n", s)
	}
}
