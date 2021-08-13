package diagrams

import (
	"reflect"
	"strings"
	"testing"
)

func TestDiagram_Add(t *testing.T) {
	d := Diagram{
		preamble:    []string{},
		includes:    []string{},
		renderables: []Renderable{},
		skinParams:  map[string]string{},
	}
	comment := NewComment("Hello comment")
	d.Add(comment)

	// assert that the renderable was added
	if !reflect.DeepEqual(d.renderables[0], comment) {
		t.Errorf("Add() Did not find the correct renderable in the renderables slice at index 0")
	}

	d.Add(comment)

	// assert that adding the same renderable twice works
	if !reflect.DeepEqual(d.renderables[1], comment) {
		t.Errorf("Add() Did not find the correct renderable in the renderables slice at index 1")
	}
}

func TestDiagram_Include(t *testing.T) {
	d := Diagram{
		preamble:    []string{},
		includes:    []string{},
		renderables: []Renderable{},
		skinParams:  map[string]string{},
	}

	d.Include("foo")

	// assert that the include was added
	if !reflect.DeepEqual(d.includes[0], "foo") {
		t.Errorf("Include() Did not find the correct value in the includes slice at index 0")
	}

	d.Include("foo")

	// assert that adding the same include twice works
	if !reflect.DeepEqual(d.includes[1], "foo") {
		t.Errorf("Include() Did not find the correct value in the includes slice at index 1")
	}
}

func TestDiagram_Preamble(t *testing.T) {
	d := Diagram{
		preamble:    []string{},
		includes:    []string{},
		renderables: []Renderable{},
		skinParams:  map[string]string{},
	}

	d.Preamble("foo")

	// assert that the preamble line was added
	if !reflect.DeepEqual(d.preamble[0], "foo") {
		t.Errorf("Preamble() Did not find the correct value in the preamble slice at index 0")
	}

	d.Preamble("foo")

	// assert that adding the same preamble line twice works
	if !reflect.DeepEqual(d.preamble[1], "foo") {
		t.Errorf("Preamble() Did not find the correct value in the preamble slice at index 1")
	}
}

func TestDiagram_Render(t *testing.T) {
	type fields struct {
		preamble    []string
		includes    []string
		renderables []Renderable
		skinParams  map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
	}{
		{
			"test preamble",
			fields{
				preamble:    []string{"foo", "bar"},
				includes:    []string{},
				renderables: []Renderable{},
				skinParams:  map[string]string{},
			},
			`
@startuml
foo
bar
@enduml
`,
		},
		{
			"test includes",
			fields{
				preamble:    []string{},
				includes:    []string{"foo", "bar"},
				renderables: []Renderable{},
				skinParams:  map[string]string{},
			},
			`
@startuml
!include foo
!include bar
@enduml
`,
		},
		{
			"test renderables",
			fields{
				preamble:    []string{},
				includes:    []string{},
				renderables: []Renderable{NewComment("Hello comment"), NewComment("Hello second comment")},
				skinParams:  map[string]string{},
			},
			`
@startuml
' Hello comment
' Hello second comment
@enduml
`,
		},
		{
			"test skinparams",
			fields{
				preamble:    []string{},
				includes:    []string{},
				renderables: []Renderable{},
				skinParams: map[string]string{
					"monochrome":  "true",
					"roundCorner": "15",
				},
			},
			`
@startuml
skinparam monochrome true
skinparam roundCorner 15
@enduml
`,
		},
		{
			"all fields",
			fields{
				preamble:    []string{"foo", "bar"},
				includes:    []string{"inc1", "inc2"},
				renderables: []Renderable{NewComment("Hello comment"), NewComment("Hello second comment")},
				skinParams: map[string]string{
					"monochrome":  "true",
					"roundCorner": "15",
				},
			},
			`
@startuml
foo
bar
skinparam monochrome true
skinparam roundCorner 15
!include inc1
!include inc2
' Hello comment
' Hello second comment
@enduml
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Diagram{
				preamble:    tt.fields.preamble,
				includes:    tt.fields.includes,
				renderables: tt.fields.renderables,
				skinParams:  tt.fields.skinParams,
			}
			writer := NewWriter()
			if err := d.Render(&writer); err != nil {
				t.Error("Render() got error but wanted none")
			}
			actual := writer.String()
			// remove the leading \n added when writing the test spec
			want := strings.TrimLeft(tt.want, "\n")
			if !reflect.DeepEqual(actual, want) {
				t.Errorf(
					"Render() string = %v, want %v",
					strings.ReplaceAll(actual, "\n", "\\n"),
					strings.ReplaceAll(want, "\n", "\\n"),
				)
			}
		})
	}
}

func TestDiagram_SkinParam(t *testing.T) {
	d := Diagram{
		preamble:    []string{},
		includes:    []string{},
		renderables: []Renderable{},
		skinParams:  map[string]string{},
	}

	d.SkinParam("key", "val1")

	// assert that the skinparam was set
	if !reflect.DeepEqual(d.skinParams["key"], "val1") {
		t.Errorf("SkinParam() Did not find the correct skin param value for 'key'")
	}

	d.SkinParam("key", "val2")

	// assert that skin params can be overwritten
	if !reflect.DeepEqual(d.skinParams["key"], "val2") {
		t.Errorf("SkinParam() Did not find the correct skin param value for 'key'")
	}
}

func TestNewDiagram(t *testing.T) {
	d := NewDiagram()

	if !reflect.DeepEqual(d.skinParams, map[string]string{}) {
		t.Error("Map skinParams not initialized correctly")
	}
	if !reflect.DeepEqual(d.includes, []string{}) {
		t.Error("Slice includes not initialized correctly")
	}
	if !reflect.DeepEqual(d.renderables, []Renderable{}) {
		t.Error("Slice renderables not initialized correctly")
	}
	if !reflect.DeepEqual(d.preamble, []string{}) {
		t.Error("Slice preamble not initialized correctly")
	}

}
