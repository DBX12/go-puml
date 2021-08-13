package diagrams

type Diagram struct {
	preamble    []string
	includes    []string
	renderables []Renderable
	skinParams  map[string]string
}

func NewDiagram() *Diagram {
	return &Diagram{
		preamble:    []string{},
		includes:    []string{},
		renderables: []Renderable{},
		skinParams:  map[string]string{},
	}
}

// Add a Renderable to the diagram
func (d *Diagram) Add(r ...Renderable) {
	d.renderables = append(d.renderables, r...)
}

// Include another file or library
func (d *Diagram) Include(inc ...string) *Diagram {
	d.includes = append(d.includes, inc...)
	return d
}

// Adds a Preamble to the diagram
func (d *Diagram) Preamble(str ...string) *Diagram {
	d.preamble = append(d.preamble, str...)
	return d
}

// Set a skin param for the diagram
func (d *Diagram) SkinParam(key string, value string) *Diagram {
	d.skinParams[key] = value
	return d
}

func (d *Diagram) Render(writer *Writer) error {
	writer.Println("@startuml")
	for _, preambleLine := range d.preamble {
		writer.Println(preambleLine)
	}
	for key, value := range d.skinParams {
		writer.Printf("skinparam %s %s\n", key, value)
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
