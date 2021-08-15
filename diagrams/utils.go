package diagrams

import (
	"errors"
	"strings"
)

type invalidCharacter struct {
	invalid     string
	replacement string
	displayName string
}

func getInvalidCharacters() []invalidCharacter {
	return []invalidCharacter{
		{" ", "", "whitespace"},
		{"-", "_", "hyphen"},
		{":", "_", "colon"},
		{"/", "_", "slash"},
	}
}

func SanitizeId(id string) string {
	var oldNew []string
	for _, definition := range getInvalidCharacters() {
		oldNew = append(oldNew, definition.invalid, definition.replacement)
	}
	replacer := strings.NewReplacer(oldNew...)
	return replacer.Replace(id)
}

func assertValidId(name string) error {
	for _, definition := range getInvalidCharacters() {
		if strings.Contains(name, definition.invalid) {
			return errors.New("id must not contain " + definition.displayName)
		}
	}

	return nil
}

//ConditionalPrintf only prints the formatted string to the writer if the given
//value is not an empty string
func ConditionalPrintf(writer *Writer, format string, value string) {
	if value != "" {
		writer.Printf(format, value)
	}
}

//renderInnerElements renders inner elements enclosed in { if the renderables
//slice slice contains at least one element
func renderInnerElements(writer *Writer, renderables []Renderable) error {
	if len(renderables) == 0 {
		return nil
	}
	writer.Println(" {")
	for _, renderable := range renderables {
		err := renderable.Render(writer)
		if err != nil {
			return err
		}
	}
	writer.Print("}")
	return writer.GetError()
}
