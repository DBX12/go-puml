package generic

import "github.com/dbx12/go-puml/diagrams"

type Link struct {
	leftId  string
	rightId string
	config  *LinkConfig
}

type LinkConfig struct {
	Line  string
	Label string
}

func NewLink(left diagrams.Linkable, right diagrams.Linkable, config *LinkConfig) *Link {
	return NewLinkFromIds(left.GetId(), right.GetId(), config)
}

func NewLinkFromIds(leftId string, rightId string, config *LinkConfig) *Link {
	if config == nil {
		config = &LinkConfig{
			Line:  "--",
			Label: "",
		}
	}
	return &Link{
		leftId:  leftId,
		rightId: rightId,
		config:  config,
	}
}

func (l Link) Render(writer *diagrams.Writer) error {
	writer.Printf("%s %s %s", l.leftId, l.config.Line, l.rightId)
	if len(l.config.Label) > 0 {
		writer.Printf(" : %s", l.config.Label)
	}
	writer.Println("")
	return writer.GetError()
}
