// Copyright 2023-2024 Takashi Takizawa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

const dataDir = "_data"

type Data struct {
	ForegroundSystemColors []Color
	BackgroundSystemColors []Color
	Colors                 []Color
}

type Color struct {
	Name             string
	HtmlHexColorCode string
	AnsiColorCode    string
}

func main() {
	colors := readColors("8bit-color.csv")
	foregroundSystemColors := readColors("foreground-system-color.csv")
	backgroundSystemColors := readColors("background-system-color.csv")
	data := Data{
		ForegroundSystemColors: foregroundSystemColors,
		BackgroundSystemColors: backgroundSystemColors,
		Colors:                 colors,
	}
	generate(data, templateCode, "tables.go")
	generate(data, templateText, "colors.md")
}

func readColors(filename string) []Color {
	var colors []Color
	f, err := os.Open(filepath.Join(dataDir, filename))
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		if len(record) == 0 {
			continue
		}
		c := Color{
			Name:             record[0],
			HtmlHexColorCode: record[1],
			AnsiColorCode:    record[2],
		}
		colors = append(colors, c)
	}
	return colors
}

func generate(data Data, text string, filename string) {
	b := bytes.NewBuffer([]byte{})
	tmpl, err := template.New("table").Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(b, data)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.CreateTemp(".", filename+".*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())

	f.Write(b.Bytes())
	f.Close()
	os.Rename(f.Name(), filename)
}

var templateCode = `// Copyright 2021-2024 Takashi Takizawa. All rights reserved.
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
{{range .ForegroundSystemColors}}
	// {{.Name}} for system color
	System{{.Name}}
{{end}}

	/*
	 * 8-bit colors
	 */
{{range .Colors}}
	// {{.Name}}: HTML Hex Color Code {{.HtmlHexColorCode}}, Ansi 8-bit Color Code {{.AnsiColorCode}}
	{{.Name}}
{{end}}
)

var colorNames = []string{
	// Default foreground color
	"Default",

	/*
	 * system colors
	 */
{{range .ForegroundSystemColors}}
	"System{{.Name}}",{{end}}

	/*
	 * 8-bit colors
	 */
{{range .Colors}}
	"{{.Name}}",{{end}}
}

var colorParameters = [...]string{
	// Default foreground color
	Default: "39",

	/*
	 * system colors
	 */
{{range .ForegroundSystemColors}}
	System{{.Name}}: "{{.AnsiColorCode}}",{{end}}

	/*
	 * 8-bit colors
	 */
{{range .Colors}}
	{{.Name}}: "38;5;{{.AnsiColorCode}}",{{end}}
}

// Background Color name
const (
	// Default background color
	DefaultBackground = BackgroundColor(iota)

	/*
	 * system colors
	 */
{{range .BackgroundSystemColors}}
	// {{.Name}} for system background color
	System{{.Name}}Background
{{end}}

	/*
	 * 8-bit colors
	 */
{{range .Colors}}
	// {{.Name}} for background color: HTML Hex Color Code {{.HtmlHexColorCode}}, Ansi 8-bit Color Code {{.AnsiColorCode}}
	{{.Name}}Background
{{end}}
)

var backgroundColorNames = []string{
	// Default background color
	"Default",

	/*
	 * system colors
	 */
{{range .BackgroundSystemColors}}
	"System{{.Name}}",{{end}}

	/*
	 * 8-bit colors
	 */
{{range .Colors}}
	"{{.Name}}",{{end}}
}

var backgroundColorParameters = [...]string{
	// Default background color
	DefaultBackground: "49",

	/*
	 * system colors
	 */
{{range .BackgroundSystemColors}}
	System{{.Name}}Background: "{{.AnsiColorCode}}",{{end}}

	/*
	 * 8-bit colors
	 */
{{range .Colors}}
	{{.Name}}Background: "48;5;{{.AnsiColorCode}}",{{end}}
}
`

var templateText = `
# Supported colors

## System Colors
{{range .ForegroundSystemColors}}
- <span style="color:{{.HtmlHexColorCode}};">System{{.Name}}</span>{{end}}

## Colors
{{range .Colors}}
- <span style="color:{{.HtmlHexColorCode}};">{{.Name}}</span>{{end}}

`
