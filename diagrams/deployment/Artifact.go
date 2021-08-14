package deployment

import "github.com/dbx12/go-puml/diagrams"

type Artifact struct {
	*diagrams.ElementWithChildren
}

func NewArtifact(id string, config *diagrams.ElementConfig) *Artifact {
	return &Artifact{
		diagrams.NewElementWithChildren("artifact", id, config),
	}
}
