package deployment

import "github.com/dbx12/go-puml/diagrams"

type Actor struct {
	*diagrams.Element
}

func NewActor(id string, config *diagrams.ElementConfig) *Actor {
	return &Actor{
		diagrams.NewElement("actor", id, config),
	}
}
