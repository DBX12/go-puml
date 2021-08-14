package deployment

import "github.com/dbx12/go-puml/diagrams"

type Database struct {
	*diagrams.ElementWithChildren
}

//goland:noinspection GoUnusedExportedFunction
func NewDatabase(id string, config *diagrams.ElementConfig) *Database {
	return &Database{
		diagrams.NewElementWithChildren("database", id, config),
	}
}
