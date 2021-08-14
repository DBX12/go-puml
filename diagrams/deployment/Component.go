package deployment

import "github.com/dbx12/go-puml/diagrams"

type Component struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewComponent(id string, config *diagrams.ElementConfig) *Component {
	return &Component{
		diagrams.NewElementWithChildren("component", id, config),
	}
}
