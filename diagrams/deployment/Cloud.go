package deployment

import "github.com/dbx12/go-puml/diagrams"

type Cloud struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewCloud(id string, config *diagrams.ElementConfig) *Cloud {
	return &Cloud{
		diagrams.NewElementWithChildren("cloud", id, config),
	}
}
