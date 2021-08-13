package diagrams

type Link struct {
	leftId  string
	rightId string
	config  *LinkConfig
}

type LinkConfig struct {
	Line  string
	Label string
}

func NewLink(left Linkable, right Linkable, config *LinkConfig) *Link {
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

func (l Link) Render(writer *Writer) error {
	writer.Printf("%s %s %s", l.leftId, l.config.Line, l.rightId)
	if len(l.config.Label) > 0 {
		writer.Printf(" : %s", l.config.Label)
	}
	writer.Println("")
	return writer.GetError()
}
