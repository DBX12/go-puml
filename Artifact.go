package diagrams

type Artifact struct {
	element element
	children
}

func NewArtifact(id string, displayName string, config *ElementConfig) *Artifact {
	return &Artifact{
		element: newElement("artifact", id, displayName, config),
	}
}

func (a Artifact) Render(writer *Writer) error {
	if err := a.element.render(writer); err != nil {
		return err
	}
	writer.Println(" {")
	if err := a.children.render(writer); err != nil {
		return err
	}
	writer.Println("}")
	return writer.GetError()
}
