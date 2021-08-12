package diagrams

type Renderable interface {
	Render(writer *Writer) error
}

type Linkable interface {
	GetId() string
}

type Container interface {
	Add(r ... Renderable)
}