package diagrams

import (
	"strings"
)

type ElementConfig struct {
	DisplayName string
	Stereotype  string
	BodyColor   string
	LineType    string
	LineColor   string
	TextColor   string
}

func (e ElementConfig) render(writer *Writer) {
	var parts []string
	if e.BodyColor != "" {
		parts = append(parts, "#"+e.BodyColor)
	}
	if e.LineColor != "" {
		parts = append(parts, "line:"+e.LineColor)
	}
	if e.LineType != "" {
		parts = append(parts, "line."+e.LineType)
	}
	if e.TextColor != "" {
		parts = append(parts, "text:"+e.TextColor)
	}
	// if we have something to write, put a space between it and the last written string (e.g. display name)
	if len(parts) > 0 {
		writer.Print(" ")
		writer.Print(strings.Join(parts, ";"))
	}
}
