package deployment

import "github.com/dbx12/go-puml/diagrams"

type Stack struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewStack(id string, config *diagrams.ElementConfig) *Stack {
	return &Stack{
		diagrams.NewElementWithChildren("stack", id, config),
	}
}
