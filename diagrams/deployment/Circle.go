package deployment

import "github.com/dbx12/go-puml/diagrams"

type Circle struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewCircle(id string, config *diagrams.ElementConfig) *Circle {
	return &Circle{
		diagrams.NewElement("circle", id, config),
	}
}
