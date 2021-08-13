package diagrams

import "strings"

type Comment struct {
	text string
}

func (c Comment) Render(writer *Writer) error {
	// remove trailing newlines
	c.text = strings.TrimRight(c.text, "\n")
	lines := strings.Split(c.text, "\n")
	for _, s := range lines {
		writer.Printf("' %s\n", s)
	}
	return writer.GetError()
}

func NewComment(text string) *Comment {
	return &Comment{text: text}
}
