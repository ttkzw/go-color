// Copyright 2021-2024 Takashi Takizawa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Code generated by running \"go generate\" in github.com/ttkzw/go-color. DO NOT EDIT.

package color

// Color name
const (
	// Default foreground color
	Default = Color(iota)

	/*
	 * system colors
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
	 * 8-bit colors
	 */

	// Red: HTML Hex Color Code #FF0000, Ansi 8-bit Color Code 196
	Red

	// Orange: HTML Hex Color Code #FF8700, Ansi 8-bit Color Code 208
	Orange

	// Yellow: HTML Hex Color Code #FFFF00, Ansi 8-bit Color Code 226
	Yellow

	// Chartreuse: HTML Hex Color Code #87FF00, Ansi 8-bit Color Code 118
	Chartreuse

	// Green: HTML Hex Color Code #00FF00, Ansi 8-bit Color Code 46
	Green

	// SpringGreen: HTML Hex Color Code #00FF87, Ansi 8-bit Color Code 48
	SpringGreen

	// Cyan: HTML Hex Color Code #00FFFF, Ansi 8-bit Color Code 51
	Cyan

	// Azure: HTML Hex Color Code #0087FF, Ansi 8-bit Color Code 33
	Azure

	// Blue: HTML Hex Color Code #0000FF, Ansi 8-bit Color Code 21
	Blue

	// Violet: HTML Hex Color Code #8700FF, Ansi 8-bit Color Code 93
	Violet

	// Magenta: HTML Hex Color Code #FF00FF, Ansi 8-bit Color Code 201
	Magenta

	// Rose: HTML Hex Color Code #FF0087, Ansi 8-bit Color Code 198
	Rose

	// White: HTML Hex Color Code #FFFFFF, Ansi 8-bit Color Code 231
	White

	// Grey: HTML Hex Color Code #808080, Ansi 8-bit Color Code 244
	Grey

	// Black: HTML Hex Color Code #000000, Ansi 8-bit Color Code 16
	Black
)

var colorNames = []string{
	// Default foreground color
	"Default",

	/*
	 * system colors
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
	 * 8-bit colors
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

var colorParameters = [...]string{
	// Default foreground color
	Default: "39",

	/*
	 * system colors
	 */

	SystemBlack:         "30",
	SystemRed:           "31",
	SystemGreen:         "32",
	SystemYellow:        "33",
	SystemBlue:          "34",
	SystemMagenta:       "35",
	SystemCyan:          "36",
	SystemWhite:         "37",
	SystemBrightBlack:   "90",
	SystemBrightRed:     "91",
	SystemBrightGreen:   "92",
	SystemBrightYellow:  "93",
	SystemBrightBlue:    "94",
	SystemBrightMagenta: "95",
	SystemBrightCyan:    "96",
	SystemBrightWhite:   "97",

	/*
	 * 8-bit colors
	 */

	Red:         "38;5;196",
	Orange:      "38;5;208",
	Yellow:      "38;5;226",
	Chartreuse:  "38;5;118",
	Green:       "38;5;46",
	SpringGreen: "38;5;48",
	Cyan:        "38;5;51",
	Azure:       "38;5;33",
	Blue:        "38;5;21",
	Violet:      "38;5;93",
	Magenta:     "38;5;201",
	Rose:        "38;5;198",
	White:       "38;5;231",
	Grey:        "38;5;244",
	Black:       "38;5;16",
}

// Background Color name
const (
	// Default background color
	DefaultBackground = BackgroundColor(iota)

	/*
	 * system colors
	 */

	// Black for system background color
	SystemBlackBackground

	// Red for system background color
	SystemRedBackground

	// Green for system background color
	SystemGreenBackground

	// Yellow for system background color
	SystemYellowBackground

	// Blue for system background color
	SystemBlueBackground

	// Magenta for system background color
	SystemMagentaBackground

	// Cyan for system background color
	SystemCyanBackground

	// White for system background color
	SystemWhiteBackground

	// BrightBlack for system background color
	SystemBrightBlackBackground

	// BrightRed for system background color
	SystemBrightRedBackground

	// BrightGreen for system background color
	SystemBrightGreenBackground

	// BrightYellow for system background color
	SystemBrightYellowBackground

	// BrightBlue for system background color
	SystemBrightBlueBackground

	// BrightMagenta for system background color
	SystemBrightMagentaBackground

	// BrightCyan for system background color
	SystemBrightCyanBackground

	// BrightWhite for system background color
	SystemBrightWhiteBackground

	/*
	 * 8-bit colors
	 */

	// Red for background color: HTML Hex Color Code #FF0000, Ansi 8-bit Color Code 196
	RedBackground

	// Orange for background color: HTML Hex Color Code #FF8700, Ansi 8-bit Color Code 208
	OrangeBackground

	// Yellow for background color: HTML Hex Color Code #FFFF00, Ansi 8-bit Color Code 226
	YellowBackground

	// Chartreuse for background color: HTML Hex Color Code #87FF00, Ansi 8-bit Color Code 118
	ChartreuseBackground

	// Green for background color: HTML Hex Color Code #00FF00, Ansi 8-bit Color Code 46
	GreenBackground

	// SpringGreen for background color: HTML Hex Color Code #00FF87, Ansi 8-bit Color Code 48
	SpringGreenBackground

	// Cyan for background color: HTML Hex Color Code #00FFFF, Ansi 8-bit Color Code 51
	CyanBackground

	// Azure for background color: HTML Hex Color Code #0087FF, Ansi 8-bit Color Code 33
	AzureBackground

	// Blue for background color: HTML Hex Color Code #0000FF, Ansi 8-bit Color Code 21
	BlueBackground

	// Violet for background color: HTML Hex Color Code #8700FF, Ansi 8-bit Color Code 93
	VioletBackground

	// Magenta for background color: HTML Hex Color Code #FF00FF, Ansi 8-bit Color Code 201
	MagentaBackground

	// Rose for background color: HTML Hex Color Code #FF0087, Ansi 8-bit Color Code 198
	RoseBackground

	// White for background color: HTML Hex Color Code #FFFFFF, Ansi 8-bit Color Code 231
	WhiteBackground

	// Grey for background color: HTML Hex Color Code #808080, Ansi 8-bit Color Code 244
	GreyBackground

	// Black for background color: HTML Hex Color Code #000000, Ansi 8-bit Color Code 16
	BlackBackground
)

var backgroundColorNames = []string{
	// Default background color
	"Default",

	/*
	 * system colors
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
	 * 8-bit colors
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

var backgroundColorParameters = [...]string{
	// Default background color
	DefaultBackground: "49",

	/*
	 * system colors
	 */

	SystemBlackBackground:         "40",
	SystemRedBackground:           "41",
	SystemGreenBackground:         "42",
	SystemYellowBackground:        "43",
	SystemBlueBackground:          "44",
	SystemMagentaBackground:       "45",
	SystemCyanBackground:          "46",
	SystemWhiteBackground:         "47",
	SystemBrightBlackBackground:   "100",
	SystemBrightRedBackground:     "101",
	SystemBrightGreenBackground:   "102",
	SystemBrightYellowBackground:  "103",
	SystemBrightBlueBackground:    "104",
	SystemBrightMagentaBackground: "105",
	SystemBrightCyanBackground:    "106",
	SystemBrightWhiteBackground:   "107",

	/*
	 * 8-bit colors
	 */

	RedBackground:         "48;5;196",
	OrangeBackground:      "48;5;208",
	YellowBackground:      "48;5;226",
	ChartreuseBackground:  "48;5;118",
	GreenBackground:       "48;5;46",
	SpringGreenBackground: "48;5;48",
	CyanBackground:        "48;5;51",
	AzureBackground:       "48;5;33",
	BlueBackground:        "48;5;21",
	VioletBackground:      "48;5;93",
	MagentaBackground:     "48;5;201",
	RoseBackground:        "48;5;198",
	WhiteBackground:       "48;5;231",
	GreyBackground:        "48;5;244",
	BlackBackground:       "48;5;16",
}
