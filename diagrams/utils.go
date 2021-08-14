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

//conditionalPrintf only prints the formatted string to the writer if the given
//value is not an empty string
func conditionalPrintf(writer *Writer, format string, value string) {
	if value != "" {
		writer.Printf(format, value)
	}
}
