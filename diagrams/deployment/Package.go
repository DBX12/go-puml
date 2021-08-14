package deployment

import "github.com/dbx12/go-puml/diagrams"

type Package struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewPackage(id string, config *diagrams.ElementConfig) *Package {
	return &Package{
		diagrams.NewElementWithChildren("package", id, config),
	}
}
