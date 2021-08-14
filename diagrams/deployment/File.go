package deployment

import "github.com/dbx12/go-puml/diagrams"

type File struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewFile(id string, config *diagrams.ElementConfig) *File {
	return &File{
		diagrams.NewElementWithChildren("file", id, config),
	}
}
