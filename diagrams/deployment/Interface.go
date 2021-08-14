package deployment

import "github.com/dbx12/go-puml/diagrams"

type Interface struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewInterface(id string, config *diagrams.ElementConfig) *Interface {
	return &Interface{
		diagrams.NewElement("interface", id, config),
	}
}
