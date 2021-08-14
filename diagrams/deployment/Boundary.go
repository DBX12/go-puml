package deployment

import "github.com/dbx12/go-puml/diagrams"

type Boundary struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewBoundary(id string, config *diagrams.ElementConfig) *Boundary {
	return &Boundary{
		diagrams.NewElement("boundary", id, config),
	}
}
