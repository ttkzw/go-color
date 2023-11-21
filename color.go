// Copyright 2021 Takashi Takizawa. All rights reserved.
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

const Version = "0.1.6"

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

func (c Color) escapeSequence() string {
	if 0 <= c && int(c) < len(colorEscapeSequences) {
		return colorEscapeSequences[c]
	}
	return colorEscapeSequences[0]
}

// Colorize colorizes the string.
func (c Color) Colorize(str string) string {
	isTerminal := true
	if terminalDetection {
		isTerminal = isatty.IsTerminal(os.Stdout.Fd())
	}
	if !isTerminal {
		return str
	}
	return fmt.Sprintf("%s%s%s", c.escapeSequence(), str, Default.escapeSequence())
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
