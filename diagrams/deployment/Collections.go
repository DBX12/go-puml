package deployment

import "github.com/dbx12/go-puml/diagrams"

type Collections struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewCollections(id string, config *diagrams.ElementConfig) *Collections {
	return &Collections{
		diagrams.NewElement("collections", id, config),
	}
}
