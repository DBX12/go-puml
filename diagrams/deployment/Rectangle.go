package deployment

import "github.com/dbx12/go-puml/diagrams"

type Rectangle struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewRectangle(id string, config *diagrams.ElementConfig) *Rectangle {
	return &Rectangle{
		diagrams.NewElementWithChildren("rectangle", id, config),
	}
}
