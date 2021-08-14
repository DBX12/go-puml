package deployment

import "github.com/dbx12/go-puml/diagrams"

type Control struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewControl(id string, config *diagrams.ElementConfig) *Control {
	return &Control{
		diagrams.NewElement("control", id, config),
	}
}
