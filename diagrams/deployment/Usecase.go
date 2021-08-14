package deployment

import "github.com/dbx12/go-puml/diagrams"

type Usecase struct {
	*diagrams.Element
}

//goland:noinspection GoUnusedExportedFunction
func NewUsecase(id string, config *diagrams.ElementConfig) *Usecase {
	return &Usecase{
		diagrams.NewElement("usecase", id, config),
	}
}
