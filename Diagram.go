package diagrams

type Diagram struct {
	preamble    []string
	includes    []string
	renderables []Renderable
}

func NewDiagram() *Diagram {
	return &Diagram{
		preamble:    []string{},
		includes:    []string{},
		renderables: []Renderable{},
	}
}

func (d *Diagram) Add(r ...Renderable) {
	d.renderables = append(d.renderables, r...)
}

func (d *Diagram) Include(inc ...string) *Diagram {
	d.includes = append(d.includes, inc...)
	return d
}

func (d *Diagram) Preamble(str ...string) *Diagram {
	d.preamble = append(d.preamble, str...)
	return d
}

func (d *Diagram) Render(writer *Writer) error {
	writer.Println("@startuml")
	for _, preambleLine := range d.preamble {
		writer.Println(preambleLine)
	}
	for _, include := range d.includes {
		writer.Printf("!include %s\n", include)
	}
	for _, renderable := range d.renderables {
		if err := renderable.Render(writer); err != nil {
			return err
		}
	}
	writer.Println("@enduml")
	return writer.GetError()
}
