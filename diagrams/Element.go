package diagrams

// Element is the base type for all elements which do not have inner elements (children)
type Element struct {
	TypeString string
	Id         string
	Config     *ElementConfig
}

func NewElement(typeString string, id string, elementConfig *ElementConfig) *Element {
	if elementConfig == nil {
		elementConfig = &ElementConfig{}
	}
	return &Element{
		TypeString: typeString,
		Id:         SanitizeId(id),
		Config:     elementConfig,
	}
}

func (e Element) Render(writer *Writer) error {
	writer.Printf("%s %s", e.TypeString, e.Id)
	ConditionalPrintf(writer, " <<%s>>", e.Config.Stereotype)
	ConditionalPrintf(writer, " as \"%s\"", e.Config.DisplayName)
	e.Config.render(writer)
	writer.Println("")
	return writer.GetError()
}

func (e Element) GetId() string {
	return e.Id
}
