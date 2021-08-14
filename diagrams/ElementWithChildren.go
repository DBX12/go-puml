package diagrams

type ElementWithChildren struct {
	TypeString string
	Id         string
	Config     *ElementConfig
	Children   []Renderable
}

func (e *ElementWithChildren) Add(r ...Renderable) {
	e.Children = append(e.Children, r...)
}

func NewElementWithChildren(typeString string, id string, elementConfig *ElementConfig) *ElementWithChildren {
	if elementConfig == nil {
		elementConfig = &ElementConfig{}
	}
	return &ElementWithChildren{
		TypeString: typeString,
		Id:         SanitizeId(id),
		Config:     elementConfig,
		Children:   []Renderable{},
	}
}

func (e ElementWithChildren) Render(writer *Writer) error {
	writer.Printf("%s %s", e.TypeString, e.Id)
	conditionalPrintf(writer, " <<%s>>", e.Config.Stereotype)
	conditionalPrintf(writer, " as \"%s\"", e.Config.DisplayName)
	e.Config.render(writer)
	if err := renderInnerElements(writer, e.Children); err != nil {
		return err
	}
	writer.Println("")
	return writer.GetError()
}

func (e ElementWithChildren) GetId() string {
	return e.Id
}
