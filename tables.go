// Copyright 2021 Takashi Takizawa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file is genarated by the script 'scripts/create-tables.sh'.

package color

// Color name
const (
	Default = Color(iota)

	/*
	 * 16 colors for system
	 */

	// Black for system color
	SystemBlack

	// Red for system color
	SystemRed

	// Green for system color
	SystemGreen

	// Yellow for system color
	SystemYellow

	// Blue for system color
	SystemBlue

	// Magenta for system color
	SystemMagenta

	// Cyan for system color
	SystemCyan

	// White for system color
	SystemWhite

	// BrightBlack for system color
	SystemBrightBlack

	// BrightRed for system color
	SystemBrightRed

	// BrightGreen for system color
	SystemBrightGreen

	// BrightYellow for system color
	SystemBrightYellow

	// BrightBlue for system color
	SystemBrightBlue

	// BrightMagenta for system color
	SystemBrightMagenta

	// BrightCyan for system color
	SystemBrightCyan

	// BrightWhite for system color
	SystemBrightWhite

	/*
	 * 256 colors
	 */

	// Red: Hex #FF0000, 256-Color Code 196 (#FF0000)
	Red

	// Orange: Hex #FF8000, 256-Color Code 208 (#FF8700)
	Orange

	// Yellow: Hex #FFFF00, 256-Color Code 226 (#FFFF00)
	Yellow

	// Chartreuse: Hex #80FF00, 256-Color Code 118 (#87FF00)
	Chartreuse

	// Green: Hex #00FF00, 256-Color Code 46 (#00FF00)
	Green

	// SpringGreen: Hex #00FF80, 256-Color Code 48 (#00FF87)
	SpringGreen

	// Cyan: Hex #00FFFF, 256-Color Code 51 (#00FFFF)
	Cyan

	// Azure: Hex #007FFF, 256-Color Code 195 (#0087FF)
	Azure

	// Blue: Hex #0000FF, 256-Color Code 21 (#0000FF)
	Blue

	// Violet: Hex #7F00FF, 256-Color Code 93 (#8700FF)
	Violet

	// Magenta: Hex #FF00FF, 256-Color Code 201 (#FF00FF)
	Magenta

	// Rose: Hex #FF0080, 256-Color Code 198 (#FF0087)
	Rose

	// White: Hex #FFFFFF, 256-Color Code 231 (#FFFFFF)
	White

	// Grey: Hex #808080, 256-Color Code 244 (#808080)
	Grey

	// Black: Hex #000000, 256-Color Code 16 (#000000)
	Black

)

var colorNames = []string{
	"Default",

	/*
	 * 16 colors for system
	 */

	"SystemBlack",
	"SystemRed",
	"SystemGreen",
	"SystemYellow",
	"SystemBlue",
	"SystemMagenta",
	"SystemCyan",
	"SystemWhite",
	"SystemBrightBlack",
	"SystemBrightRed",
	"SystemBrightGreen",
	"SystemBrightYellow",
	"SystemBrightBlue",
	"SystemBrightMagenta",
	"SystemBrightCyan",
	"SystemBrightWhite",

	/*
	 * 256 colors
	 */

	"Red",
	"Orange",
	"Yellow",
	"Chartreuse",
	"Green",
	"SpringGreen",
	"Cyan",
	"Azure",
	"Blue",
	"Violet",
	"Magenta",
	"Rose",
	"White",
	"Grey",
	"Black",
}

var colorEscapeSequences = [...]string{
	Default: "\x1b[39m",

	/*
	 * 16 colors for system
	 */

	SystemBlack:         "\x1b[30m",
	SystemRed:           "\x1b[31m",
	SystemGreen:         "\x1b[32m",
	SystemYellow:        "\x1b[33m",
	SystemBlue:          "\x1b[34m",
	SystemMagenta:       "\x1b[35m",
	SystemCyan:          "\x1b[36m",
	SystemWhite:         "\x1b[37m",
	SystemBrightBlack:   "\x1b[90m",
	SystemBrightRed:     "\x1b[91m",
	SystemBrightGreen:   "\x1b[92m",
	SystemBrightYellow:  "\x1b[93m",
	SystemBrightBlue:    "\x1b[94m",
	SystemBrightMagenta: "\x1b[95m",
	SystemBrightCyan:    "\x1b[96m",
	SystemBrightWhite:   "\x1b[97m",

	/*
	 * 256 colors
	 */

	Red:         "\x1b[38;5;196m",
	Orange:      "\x1b[38;5;208m",
	Yellow:      "\x1b[38;5;226m",
	Chartreuse:  "\x1b[38;5;118m",
	Green:       "\x1b[38;5;46m",
	SpringGreen: "\x1b[38;5;48m",
	Cyan:        "\x1b[38;5;51m",
	Azure:       "\x1b[38;5;195m",
	Blue:        "\x1b[38;5;21m",
	Violet:      "\x1b[38;5;93m",
	Magenta:     "\x1b[38;5;201m",
	Rose:        "\x1b[38;5;198m",
	White:       "\x1b[38;5;231m",
	Grey:        "\x1b[38;5;244m",
	Black:       "\x1b[38;5;16m",
}
