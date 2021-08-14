package deployment

import "github.com/dbx12/go-puml/diagrams"

type Agent struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewAgent(id string, config *diagrams.ElementConfig) *Agent {
	return &Agent{
		diagrams.NewElement("agent", id, config),
	}
}
