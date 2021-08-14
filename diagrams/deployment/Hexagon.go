package deployment

import "github.com/dbx12/go-puml/diagrams"

type Hexagon struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewHexagon(id string, config *diagrams.ElementConfig) *Hexagon {
	return &Hexagon{
		diagrams.NewElementWithChildren("hexagon", id, config),
	}
}
