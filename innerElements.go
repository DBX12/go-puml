package diagrams

type children struct {
	elements []Renderable
}

//Renders all child elements
func (c *children) render(writer *Writer) error {
	for _, child := range c.elements {
		if err := child.Render(writer); err != nil {
			return err
		}
	}
	return nil
}

//Add another Renderable element to the collection
func (c *children) Add(r ...Renderable) {
	c.elements = append(c.elements, r...)
}