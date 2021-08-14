package diagrams

// A Renderable is an element which can be rendered with a writer.
// If the element has an unique identifier and can be linked with a Link
// element, it should implement the Linkable interface.
// If the element can contain other Renderable elements, it should implement the
// Container interface.
type Renderable interface {
	// Render adds the elements textual representation to the writer's buffer
	Render(writer *Writer) error
}

// A Linkable can be linked with another Linkable by using a Link element.
// Most elements implementing Renderable implement Linkable as well.
type Linkable interface {
	Renderable
	// GetId returns the identifier of the element
	GetId() string
}

// A Container contains other Renderable elements.
type Container interface {
	Renderable
	// Adds one or more Renderable elements to the Container's children list
	Add(r ...Renderable)
}
