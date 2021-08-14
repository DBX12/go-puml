package deployment

import "github.com/dbx12/go-puml/diagrams"

type Label struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewLabel(id string, config *diagrams.ElementConfig) *Label {
	return &Label{
		diagrams.NewElement("label", id, config),
	}
}
