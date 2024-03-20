// Copyright 2021-2024 Takashi Takizawa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen_tables.go
//go:generate gofmt -w tables.go

package color

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

const Version = "0.4.0"

const (
	// csi is CONTROL SEQUENCE INTRODUCER (CSI), "ESC [".
	csi = "\x1b["

	// sgr_end_char is the end character of SGR (SELECT GRAPHIC RENDITION), which is "m".
	sgr_end_char = "m"

	// sgr_separeter is the separator of SGR (SELECT GRAPHIC RENDITION), which is ";".
	sgr_separator = ";"
)

var terminalDetection = true

// Color is a color for output.
type Color int

// NewColor creates the Color.
func NewColor(name string) (Color, error) {
	for i, v := range colorNames {
		if strings.EqualFold(v, name) {
			return Color(i), nil
		}
	}
	return 0, fmt.Errorf("unknown color name: %s", name)
}

// String returns the name of the color.
func (c Color) String() string {
	return colorNames[c]
}

func (c Color) parameter() string {
	var parameter = colorParameters[0]
	if 0 <= c && int(c) < len(colorParameters) {
		parameter = colorParameters[c]
	}
	return parameter
}

func (c Color) escapeSequence() string {
	return strings.Join([]string{csi, c.parameter(), sgr_end_char}, "")
}

// Colorize colors the string with the foreground color.
func (c Color) Colorize(str string) string {
	if terminalDetection && !isatty.IsTerminal(os.Stdout.Fd()) {
		return str
	}
	return fmt.Sprint(strings.Join([]string{c.escapeSequence(), str, Default.escapeSequence()}, ""))
}

// BackgroundColor is a background color for output.
type BackgroundColor int

// NewBackgroundColor creates the BackgroundColor.
func NewBackgroundColor(name string) (BackgroundColor, error) {
	for i, v := range backgroundColorNames {
		if strings.EqualFold(v, name) {
			return BackgroundColor(i), nil
		}
	}
	return 0, fmt.Errorf("unknown color name: %s", name)
}

// String returns the name of the color.
func (c BackgroundColor) String() string {
	return backgroundColorNames[c]
}

func (c BackgroundColor) parameter() string {
	var parameter = backgroundColorParameters[0]
	if 0 <= c && int(c) < len(backgroundColorParameters) {
		parameter = backgroundColorParameters[c]
	}
	return parameter
}

func (c BackgroundColor) escapeSequence() string {
	return strings.Join([]string{csi, c.parameter(), sgr_end_char}, "")
}

// Colorize colors the string with the background color.
func (c BackgroundColor) Colorize(str string) string {
	if terminalDetection && !isatty.IsTerminal(os.Stdout.Fd()) {
		return str
	}
	return fmt.Sprint(strings.Join([]string{c.escapeSequence(), str, DefaultBackground.escapeSequence()}, ""))
}

func SupportedColors() []Color {
	var colors []Color
	for _, name := range colorNames {
		c, _ := NewColor(name)
		if c == Default {
			continue
		}
		colors = append(colors, c)
	}
	return colors
}

// ColorizeForeground colors the string with the foreground color.
func ColorizeForeground(str string, foreground Color) string {
	return foreground.Colorize(str)
}

// ColorizeBackground colors the string with the background color.
func ColorizeBackground(str string, background BackgroundColor) string {
	return background.Colorize(str)
}

// Colorize colors the string　with foreground and background colors.
func Colorize(str string, foreground Color, background BackgroundColor) string {
	if terminalDetection && !isatty.IsTerminal(os.Stdout.Fd()) {
		return str
	}
	return strings.Join([]string{
		csi, foreground.parameter(), sgr_separator, background.parameter(), sgr_end_char,
		str,
		csi, Default.parameter(), sgr_separator, DefaultBackground.parameter(), sgr_end_char}, "")
}

// EnableTerminalDetection enables terminal detection.
// If no terminal is detected, colorization is suppressed.
// By default, the terminal detection is enabled.
func EnableTerminalDetection() {
	terminalDetection = true
}

// DisableTerminalDetection disables terminal detection.
// Even if the terminal is not detected, colorization is not suppressed.
// By default, the terminal detection is enabled.
func DisableTerminalDetection() {
	terminalDetection = false
}
