package deployment

import "github.com/dbx12/go-puml/diagrams"

type Storage struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewStorage(id string, config *diagrams.ElementConfig) *Storage {
	return &Storage{
		diagrams.NewElementWithChildren("storage", id, config),
	}
}
