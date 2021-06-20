package color_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/ttkzw/go-color"
)

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

func TestString(t *testing.T) {
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

	// System Colors
	f, _ := os.Open("testdata/system-color.csv")
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	for _, record := range records {
		c, _ := color.NewColor("System" + record[0])
		got := c.String()
		want := "System" + record[0]
		if got != want {
			t.Errorf("color.%s.String() = %s; want %s", record[0], got, want)
		}
	}

	// Colors
	f, _ = os.Open("testdata/color.csv")
	r = csv.NewReader(f)
	records, _ = r.ReadAll()
	for _, record := range records {
		c, _ := color.NewColor(record[0])
		got := c.String()
		want := record[0]
		if got != want {
			t.Errorf("color.%s.String() = %s; want %s", record[0], got, want)
		}
	}

}

func TestColorize(t *testing.T) {
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

	// System Colors
	f, _ := os.Open("testdata/system-color.csv")
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	for _, record := range records {
		c, _ := color.NewColor("System" + record[0])
		got := c.Colorize("TEXT")
		want := fmt.Sprintf("\x1b[%smTEXT\x1b[39m", record[2])
		if got != want {
			t.Errorf("color.System%s.Colorize(\"TEXT\") = %s; want %s", record[0], got, want)
		}
	}

	// Colors
	f, _ = os.Open("testdata/color.csv")
	r = csv.NewReader(f)
	records, _ = r.ReadAll()
	for _, record := range records {
		c, _ := color.NewColor(record[0])
		got := c.Colorize("TEXT")
		want := fmt.Sprintf("\x1b[38;5;%smTEXT\x1b[39m", record[2])
		if got != want {
			t.Errorf("color.%s.Colorize(\"TEXT\") = %s; want %s", record[0], got, want)
		}
	}
}
