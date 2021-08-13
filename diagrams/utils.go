package diagrams

import (
	"errors"
	"strings"
)

func MakeValidId(id string) string {
	replacer := strings.NewReplacer(
		" ", "",
		"-", "_",
	)
	return replacer.Replace(id)
}

func assertValidId(name string) error {
	invalidChars := map[string]string{
		" ": "whitespace",
		"-": "hyphen",
	}

	for char, display := range invalidChars {
		if strings.Contains(name, char) {
			return errors.New("id must not contain " + display)
		}
	}

	return nil
}
