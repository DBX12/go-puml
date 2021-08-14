package diagrams

import "strings"

// Comment renders a diagram code comment starting with a '
// The Comment.text will be invisible on the generated diagram
// Not to be confused with the Note element which results in a visible note on
// the generated diagram.
type Comment struct {
	text string
}

func NewComment(text string) *Comment {
	return &Comment{text: text}
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
