package diagrams


type Node struct {
	element element
	contents []string
}

func NewNode(id string, displayName string, contents []string) *Node {
	if err := assertValidId(id); err != nil{
		panic(err)
	}
	return &Node{
		element: element{
			elementType: "node",
			id: id,
			displayName: displayName,
		},
		contents: contents,
	}
}

func (n Node) GetId() string {
	return n.element.id
}

func (n Node) Render(writer *Writer) error {
	if err := n.element.render(writer); err != nil {
		return err
	}

	if len(n.contents) > 0 {
		writer.Println(" [")
		for _, content := range n.contents {
			writer.Println(content)
		}
		writer.Println("]")
	}else{
		writer.Println("")
	}
	return writer.GetError()
}
