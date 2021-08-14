package deployment

import "github.com/dbx12/go-puml/diagrams"

type Folder struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewFolder(id string, config *diagrams.ElementConfig) *Folder {
	return &Folder{
		diagrams.NewElementWithChildren("folder", id, config),
	}
}
