package diagrams

type File struct {
	element element
	children
}

func NewFile(id string, displayName string, config *ElementConfig) *File {
	return &File{
		element: newElement("file", id, displayName, config),
	}
}

func (f *File) GetId() string {
	return f.element.id
}

func (f File) Render(writer *Writer) error {
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
