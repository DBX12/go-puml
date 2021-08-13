package diagrams

import "strings"

type element struct {
	elementType string
	id          string
	displayName string
	config      ElementConfig
}

type ElementConfig struct {
	BodyColor string
	LineType  string
	LineColor string
	TextColor string
}

func newElement(elementType string, id string, displayName string, config *ElementConfig) element {
	if err := assertValidId(id); err != nil {
		panic(err)
	}
	if config == nil {
		config = &ElementConfig{}
	}
	return element{
		elementType: elementType,
		id:          id,
		displayName: displayName,
		config:      *config,
	}
}

func (e element) GetId() string {
	return e.id
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
	writer.Print(strings.Join(parts, ";"))
}

func (e element) render(writer *Writer) error {
	writer.Printf("%s %s", e.elementType, e.id)
	if e.displayName != "" {
		writer.Printf(" as \"%s\" ", e.displayName)
	}
	e.config.render(writer)
	return writer.GetError()
}
