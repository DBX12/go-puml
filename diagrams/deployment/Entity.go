package deployment

import "github.com/dbx12/go-puml/diagrams"

type Entity struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewEntity(id string, config *diagrams.ElementConfig) *Entity {
	return &Entity{
		diagrams.NewElement("entity", id, config),
	}
}
