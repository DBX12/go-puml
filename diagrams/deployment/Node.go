package deployment

import "github.com/dbx12/go-puml/diagrams"

type Node struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewNode(id string, config *diagrams.ElementConfig) *Node {
	return &Node{
		diagrams.NewElementWithChildren("node", id, config),
	}
}
