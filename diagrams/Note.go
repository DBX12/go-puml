package diagrams

import "strings"

type NotePosition int

const (
	BELOW NotePosition = iota
	RIGHT
	ABOVE
	LEFT
)

type Note struct {
	linkedTo string
	position NotePosition
	contents []string
}

func NewNoteForLinkable(target Linkable, position NotePosition, contents []string) *Note {
	return NewNoteForId(target.GetId(), position, contents)
}

func NewNoteForId(id string, position NotePosition, contents []string) *Note {
	return &Note{
		linkedTo: id,
		position: position,
		contents: contents,
	}
}

func (n Note) Render(writer *Writer) error {
	positions := map[NotePosition]string{
		BELOW: "bottom",
		RIGHT: "right",
		ABOVE: "top",
		LEFT:  "left",
	}
	writer.Printf("note %s of %s\n", positions[n.position], n.linkedTo)
	writer.Print(strings.Join(n.contents, "\n"))
	writer.Printf("\nendnote\n")
	return writer.GetError()
}
