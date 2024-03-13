package color_test

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/ttkzw/go-color"
)

const dataDir = "_data"

var (
	foregroundSystemColorRecords [][]string
	backgroundSystemColorRecords [][]string
	colorRecords                 [][]string
)

func init() {
	f, err := os.Open(filepath.Join(dataDir, "foreground-system-color.csv"))
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	foregroundSystemColorRecords, err = r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Open(filepath.Join(dataDir, "background-system-color.csv"))
	if err != nil {
		log.Fatal(err)
	}
	r = csv.NewReader(f)
	backgroundSystemColorRecords, err = r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Open(filepath.Join(dataDir, "8bit-color.csv"))
	if err != nil {
		log.Fatal(err)
	}
	r = csv.NewReader(f)
	colorRecords, err = r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
}

type foregroundTest struct {
	color  color.Color
	name   string
	output string
}

var foregroundTests = []foregroundTest{
	{color.SystemBrightRed, "SystemBrightRed", "\x1b[91mTEXT\x1b[39m"},
	{color.Red, "Red", "\x1b[38;5;196mTEXT\x1b[39m"},
	{color.Black, "Black", "\x1b[38;5;16mTEXT\x1b[39m"},
}

func TestNewColor(t *testing.T) {
	if _, err := color.NewColor("Unknown"); err == nil {
		t.Errorf("color.NewColor(\"Unknown\"); want error")
	}

	for _, tt := range foregroundTests {
		if got, err := color.NewColor(tt.name); got != tt.color {
			if err != nil {
				t.Fatal(err)
			}
			t.Errorf("color.NewColor(\"%s\") = %s; want %s", tt.name, got.String(), tt.name)
		}
	}
}

func TestColorString(t *testing.T) {
	for _, tt := range foregroundTests {
		if got := tt.color.String(); got != tt.name {
			t.Errorf("color.%s.String() = %s; want %s", tt.name, got, tt.name)
		}
	}

	// Foreground System Colors
	for _, record := range foregroundSystemColorRecords {
		c, _ := color.NewColor("System" + record[0])
		got := c.String()
		want := "System" + record[0]
		if got != want {
			t.Errorf("color.%s.String() = %s; want %s", record[0], got, want)
		}
	}

	// Foreground Colors
	for _, record := range colorRecords {
		c, _ := color.NewColor(record[0])
		got := c.String()
		want := record[0]
		if got != want {
			t.Errorf("color.%s.String() = %s; want %s", record[0], got, want)
		}
	}
}

func TestColorColorize(t *testing.T) {
	// Non Terminal
	color.EnableTerminalDetection()

	for _, tt := range foregroundTests {
		if got := tt.color.Colorize("TEXT"); got != "TEXT" {
			t.Errorf("color.%s.Colorize(\"TEXT\") = %s; want TEXT", tt.name, got)
		}
	}

	// Terminal
	color.DisableTerminalDetection()

	for _, tt := range foregroundTests {
		if got := tt.color.Colorize("TEXT"); got != tt.output {
			t.Errorf("color.%s.Colorize(\"TEXT\") = %s; want %s", tt.name, got, tt.output)
		}
	}

	// Foreground System Colors
	for _, record := range foregroundSystemColorRecords {
		c, _ := color.NewColor("System" + record[0])
		got := c.Colorize("TEXT")
		want := fmt.Sprintf("\x1b[%smTEXT\x1b[39m", record[2])
		if got != want {
			t.Errorf("color.System%s.Colorize(\"TEXT\") = %s; want %s", record[0], got, want)
		}
	}

	// Foreground Colors
	for _, record := range colorRecords {
		c, _ := color.NewColor(record[0])
		got := c.Colorize("TEXT")
		want := fmt.Sprintf("\x1b[38;5;%smTEXT\x1b[39m", record[2])
		if got != want {
			t.Errorf("color.%s.Colorize(\"TEXT\") = %s; want %s", record[0], got, want)
		}
	}
}

type backgroundTest struct {
	color  color.BackgroundColor
	name   string
	output string
}

var backgroundTests = []backgroundTest{
	{color.SystemBrightRedBackground, "SystemBrightRed", "\x1b[101mTEXT\x1b[49m"},
	{color.RedBackground, "Red", "\x1b[48;5;196mTEXT\x1b[49m"},
	{color.BlackBackground, "Black", "\x1b[48;5;16mTEXT\x1b[49m"},
}

func TestNewBackgroundColor(t *testing.T) {
	if _, err := color.NewBackgroundColor("Unknown"); err == nil {
		t.Errorf("color.NewBackgroundColor(\"Unknown\"); want error")
	}

	for _, tt := range backgroundTests {
		if got, err := color.NewBackgroundColor(tt.name); got != tt.color {
			if err != nil {
				t.Fatal(err)
			}
			t.Errorf("color.NewBackgroundColor(\"%s\") = %s; want %s", tt.name, got.String(), tt.name)
		}
	}
}

func TestBackgroundColorString(t *testing.T) {
	for _, tt := range backgroundTests {
		if got := tt.color.String(); got != tt.name {
			t.Errorf("color.%sBackground.String() = %s; want %s", tt.name, got, tt.name)
		}
	}

	// Background System Colors
	for _, record := range backgroundSystemColorRecords {
		c, _ := color.NewBackgroundColor("System" + record[0])
		got := c.String()
		want := "System" + record[0]
		if got != want {
			t.Errorf("color.%sBackground.String() = %s; want %s", record[0], got, want)
		}
	}

	// Background Colors
	for _, record := range colorRecords {
		c, _ := color.NewBackgroundColor(record[0])
		got := c.String()
		want := record[0]
		if got != want {
			t.Errorf("color.%sBackground.String() = %s; want %s", record[0], got, want)
		}
	}
}

func TestBackgroundColorColorize(t *testing.T) {
	// Non Terminal
	color.EnableTerminalDetection()

	for _, tt := range backgroundTests {
		if got := tt.color.Colorize("TEXT"); got != "TEXT" {
			t.Errorf("color.%sBackground.Colorize(\"TEXT\") = %s; want TEXT", tt.name, got)
		}
	}

	// Terminal
	color.DisableTerminalDetection()

	for _, tt := range backgroundTests {
		if got := tt.color.Colorize("TEXT"); got != tt.output {
			t.Errorf("color.%sBackground.Colorize(\"TEXT\") = %s; want %s", tt.name, got, tt.output)
		}
	}

	// Background System Colors
	for _, record := range backgroundSystemColorRecords {
		c, _ := color.NewBackgroundColor("System" + record[0])
		got := c.Colorize("TEXT")
		want := fmt.Sprintf("\x1b[%smTEXT\x1b[49m", record[2])
		if got != want {
			t.Errorf("color.System%sBackground.Colorize(\"TEXT\") = %s; want %s", record[0], got, want)
		}
	}

	// Background Colors
	for _, record := range colorRecords {
		c, _ := color.NewBackgroundColor(record[0])
		got := c.Colorize("TEXT")
		want := fmt.Sprintf("\x1b[48;5;%smTEXT\x1b[49m", record[2])
		if got != want {
			t.Errorf("color.%sBackground.Colorize(\"TEXT\") = %s; want %s", record[0], got, want)
		}
	}
}

func TestColorizeForeground(t *testing.T) {
	// Non Terminal
	color.EnableTerminalDetection()

	for _, tt := range foregroundTests {
		if got := color.ColorizeForeground("TEXT", tt.color); got != "TEXT" {
			t.Errorf("color.ColorizeForeground(\"TEXT\", color.%s) = %s; want TEXT", tt.name, got)
		}
	}

	// Terminal
	color.DisableTerminalDetection()

	for _, tt := range foregroundTests {
		if got := color.ColorizeForeground("TEXT", tt.color); got != tt.output {
			t.Errorf("color.ColorizeForeground(\"TEXT\", color.%s) = %s; want %s", tt.name, got, tt.output)
		}
	}

	// Foreground System Colors
	for _, record := range foregroundSystemColorRecords {
		c, _ := color.NewColor("System" + record[0])
		got := color.ColorizeForeground("TEXT", c)
		want := fmt.Sprintf("\x1b[%smTEXT\x1b[39m", record[2])
		if got != want {
			t.Errorf("color.Colorize(\"TEXT\", color.System%s = %s; want %s", record[0], got, want)
		}
	}

	// Foreground Colors
	for _, record := range colorRecords {
		c, _ := color.NewColor(record[0])
		got := color.ColorizeForeground("TEXT", c)
		want := fmt.Sprintf("\x1b[38;5;%smTEXT\x1b[39m", record[2])
		if got != want {
			t.Errorf("color.Colorize(\"TEXT\", color.%s = %s; want %s", record[0], got, want)
		}
	}
}

func TestColorizeBackground(t *testing.T) {
	// Non Terminal
	color.EnableTerminalDetection()

	for _, tt := range backgroundTests {
		if got := color.ColorizeBackground("TEXT", tt.color); got != "TEXT" {
			t.Errorf("color.ColorizeBackground(\"TEXT\", color.%sBackground) = %s; want TEXT", tt.name, got)
		}
	}

	// Terminal
	color.DisableTerminalDetection()

	for _, tt := range backgroundTests {
		if got := color.ColorizeBackground("TEXT", tt.color); got != tt.output {
			t.Errorf("color.ColorizeBackground(\"TEXT\", color.%sBackground) = %s; want %s", tt.name, got, tt.output)
		}
	}

	// Background System Colors
	for _, record := range backgroundSystemColorRecords {
		c, _ := color.NewBackgroundColor("System" + record[0])
		got := color.ColorizeBackground("TEXT", c)
		want := fmt.Sprintf("\x1b[%smTEXT\x1b[49m", record[2])
		if got != want {
			t.Errorf("color.ColorizeBackground(\"TEXT\", color.System%sBackground) = %s; want %s", record[0], got, want)
		}
	}

	// Background Colors
	for _, record := range colorRecords {
		c, _ := color.NewBackgroundColor(record[0])
		got := color.ColorizeBackground("TEXT", c)
		want := fmt.Sprintf("\x1b[48;5;%smTEXT\x1b[49m", record[2])
		if got != want {
			t.Errorf("color.ColorizeBackground(\"TEXT\", color.%sBackground) = %s; want %s", record[0], got, want)
		}
	}
}

type colorizeTest struct {
	fgColor color.Color
	fgName  string
	bgColor color.BackgroundColor
	bgName  string
	output  string
}

var colorizeTests = []colorizeTest{
	{color.SystemWhite, "SystemWhite", color.SystemBrightRedBackground, "SystemBrightRed", "\x1b[101m\x1b[37mTEXT\x1b[39m\x1b[49m"},
	{color.White, "White", color.RedBackground, "Red", "\x1b[48;5;196m\x1b[38;5;231mTEXT\x1b[39m\x1b[49m"},
	{color.White, "White", color.BlackBackground, "Black", "\x1b[48;5;16m\x1b[38;5;231mTEXT\x1b[39m\x1b[49m"},
}

func TestColorize(t *testing.T) {
	// Non Terminal
	color.EnableTerminalDetection()

	for _, tt := range colorizeTests {
		if got := color.Colorize("TEXT", tt.fgColor, tt.bgColor); got != "TEXT" {
			t.Errorf("color.Colorize(\"TEXT\", color.%s, color.%sBackground) = %s; want TEXT", tt.fgName, tt.bgName, got)
		}
	}

	// Terminal
	color.DisableTerminalDetection()

	for _, tt := range colorizeTests {
		if got := color.Colorize("TEXT", tt.fgColor, tt.bgColor); got != tt.output {
			t.Errorf("color.Colorize(\"TEXT\", color.%s, color.%sBackground) = %s; want %s", tt.fgName, tt.bgName, got, tt.output)
		}
	}
}
