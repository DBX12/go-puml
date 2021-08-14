package deployment

import "github.com/dbx12/go-puml/diagrams"

type Queue struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewQueue(id string, config *diagrams.ElementConfig) *Queue {
	return &Queue{
		diagrams.NewElementWithChildren("queue", id, config),
	}
}
