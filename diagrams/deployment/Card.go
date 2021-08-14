package deployment

import "github.com/dbx12/go-puml/diagrams"

type Card struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewCard(id string, config *diagrams.ElementConfig) *Card {
	return &Card{
		diagrams.NewElementWithChildren("card", id, config),
	}
}
