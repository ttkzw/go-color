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

func TestNewColor(t *testing.T) {
	// Unknown
	if _, err := color.NewColor("Unknown"); err == nil {
		t.Errorf("color.NewColor(\"Unknown\"); want error")
	}

	// SystemBrightRed
	if got, _ := color.NewColor("SystemBrightRed"); got != color.SystemBrightRed {
		t.Errorf("color.NewColor(\"SystemBrightRed\") = %s; want SystemBrightRed", got.String())
	}

	// Red
	if got, _ := color.NewColor("Red"); got != color.Red {
		t.Errorf("color.NewColor(\"Red\") = %s; want Red", got.String())
	}

	// Black
	if got, _ := color.NewColor("Black"); got != color.Black {
		t.Errorf("color.NewColor(\"Black\") = %s; want Black", got.String())
	}
}

func TestColorString(t *testing.T) {
	// SystemBrightRed
	if got := color.SystemBrightRed.String(); got != "SystemBrightRed" {
		t.Errorf("color.SystemBrightRed.String() = %s; want SystemBrightRed", got)
	}

	// Red
	if got := color.Red.String(); got != "Red" {
		t.Errorf("color.Red.String() = %s; want Red", got)
	}

	// Black
	if got := color.Black.String(); got != "Black" {
		t.Errorf("color.Black.String() = %s; want Black", got)
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

	// BrightRed
	if got := color.SystemBrightRed.Colorize("TEXT"); got != "TEXT" {
		t.Errorf("color.SystemBrightRed.Colorize(\"TEXT\") = %s; want TEXT", got)
	}

	// Red
	if got := color.Red.Colorize("TEXT"); got != "TEXT" {
		t.Errorf("color.Red.Colorize(\"TEXT\") = %s; want TEXT", got)
	}

	// Black
	if got := color.Black.Colorize("TEXT"); got != "TEXT" {
		t.Errorf("color.Black.Colorize(\"TEXT\") = %s; want TEXT", got)
	}

	// Terminal
	color.DisableTerminalDetection()

	// SystemBrightRed
	if got := color.SystemBrightRed.Colorize("TEXT"); got != "\x1b[91mTEXT\x1b[39m" {
		t.Errorf("color.SystemBrightRed.Colorize(\"TEXT\") = %s; want \x1b[91mTEXT\x1b[39m", got)
	}

	// Red
	if got := color.Red.Colorize("TEXT"); got != "\x1b[38;5;196mTEXT\x1b[39m" {
		t.Errorf("color.Red.Colorize(\"TEXT\") = %s; want \x1b[38;5;196mTEXT\x1b[39m", got)
	}

	// Black
	if got := color.Black.Colorize("TEXT"); got != "\x1b[38;5;16mTEXT\x1b[39m" {
		t.Errorf("color.Black.Colorize(\"TEXT\") = %s; want \x1b[38;5;16mTEXT\x1b[39m", got)
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

func TestNewBackgroundColor(t *testing.T) {
	// Unknown
	if _, err := color.NewBackgroundColor("Unknown"); err == nil {
		t.Errorf("color.NewBackgroundColor(\"Unknown\"); want error")
	}

	// SystemBrightRed
	if got, err := color.NewBackgroundColor("SystemBrightRed"); got != color.SystemBrightRedBackground {
		if err != nil {
			t.Fatal(err)
		}
		t.Errorf("color.NewBackgroundColor(\"SystemBrightRed\") = %s; want SystemBrightRedBackground", got.String())
	}

	// Red
	if got, err := color.NewBackgroundColor("Red"); got != color.RedBackground {
		if err != nil {
			t.Fatal(err)
		}
		t.Errorf("color.NewBackgroundColor(\"Red\") = %s; want RedBackground", got.String())
	}

	// Black
	if got, err := color.NewBackgroundColor("Black"); got != color.BlackBackground {
		if err != nil {
			t.Fatal(err)
		}
		t.Errorf("color.NewBackgroundColor(\"Black\") = %s; want BlackBackground", got.String())
	}
}

func TestBackgroundColorString(t *testing.T) {
	// SystemBrightRed
	if got := color.SystemBrightRedBackground.String(); got != "SystemBrightRed" {
		t.Errorf("color.SystemBrightRedBackground.String() = %s; want SystemBrightRed", got)
	}

	// Red
	if got := color.RedBackground.String(); got != "Red" {
		t.Errorf("color.RedBackground.String() = %s; want Red", got)
	}

	// Black
	if got := color.BlackBackground.String(); got != "Black" {
		t.Errorf("color.BlackBackground.String() = %s; want Black", got)
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

	// SystemBrightRed
	if got := color.SystemBrightRedBackground.Colorize("TEXT"); got != "TEXT" {
		t.Errorf("color.SystemBrightRedBackground.Colorize(\"TEXT\") = %s; want TEXT", got)
	}

	// Red
	if got := color.RedBackground.Colorize("TEXT"); got != "TEXT" {
		t.Errorf("color.RedBackground.Colorize(\"TEXT\") = %s; want TEXT", got)
	}

	// Black
	if got := color.BlackBackground.Colorize("TEXT"); got != "TEXT" {
		t.Errorf("color.BlackBackground.Colorize(\"TEXT\") = %s; want TEXT", got)
	}

	// Terminal
	color.DisableTerminalDetection()

	// SystemBrightRed
	if got := color.SystemBrightRedBackground.Colorize("TEXT"); got != "\x1b[101mTEXT\x1b[49m" {
		t.Errorf("color.SystemBrightRedBackground.Colorize(\"TEXT\") = %s; want \x1b[101mTEXT\x1b[49m", got)
	}

	// Red
	if got := color.RedBackground.Colorize("TEXT"); got != "\x1b[48;5;196mTEXT\x1b[49m" {
		t.Errorf("color.RedBackground.Colorize(\"TEXT\") = %s; want \x1b[48;5;196mTEXT\x1b[49m", got)
	}

	// Black
	if got := color.BlackBackground.Colorize("TEXT"); got != "\x1b[48;5;16mTEXT\x1b[49m" {
		t.Errorf("color.BlackBackground.Colorize(\"TEXT\") = %s; want \x1b[48;5;16mTEXT\x1b[49m", got)
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
			t.Errorf("color.%s.Colorize(\"TEXT\") = %s; want %s", record[0], got, want)
		}
	}
}
