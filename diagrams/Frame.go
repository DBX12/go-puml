package diagrams

type Frame struct {
	element element
	children
}

func (f *Frame) GetId() string {
	return f.element.id
}

func NewFrame(id string, displayName string, config *ElementConfig) *Frame {
	return &Frame{
		element: newElement("frame", id, displayName, config),
	}
}

func (f *Frame) Render(writer *Writer) error {
	if err := f.element.render(writer); err != nil {
		return err
	}
	writer.Println(" {")
	if err := f.children.render(writer); err != nil {
		return err
	}
	writer.Println("}")
	return writer.GetError()
}
