package deployment

import "github.com/dbx12/go-puml/diagrams"

type Frame struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewFrame(id string, config *diagrams.ElementConfig) *Frame {
	return &Frame{
		diagrams.NewElementWithChildren("frame", id, config),
	}
}
