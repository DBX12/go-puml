package diagrams

import (
	"reflect"
	"testing"
)

func TestElementWithChildren_GetId(t *testing.T) {
	want := "the_id"
	elementWithChildren := NewElementWithChildren("dummy", want, nil)
	got := elementWithChildren.GetId()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetId() = %v, want %v", got, want)
	}
}

func TestElementWithChildren_Render(t *testing.T) {
	filledElement := NewElementWithChildren("dummy", "the_id", &ElementConfig{DisplayName: "displayName", Stereotype: "stereo"})
	filledElement.Add(DummyRenderable{"Dummy Child"})
	tests := []struct {
		name    string
		element *ElementWithChildren
		want    string
	}{
		{
			"without displayname",
			NewElementWithChildren("dummy", "the_id", nil),
			"dummy the_id\n",
		},
		{
			"with displayname",
			NewElementWithChildren("dummy", "the_id", &ElementConfig{DisplayName: "Line 1\nLine 2"}),
			"dummy the_id as \"Line 1\nLine 2\"\n",
		},
		{
			"with stereotype",
			NewElementWithChildren("dummy", "the_id", &ElementConfig{DisplayName: "displayName", Stereotype: "stereo"}),
			"dummy the_id <<stereo>> as \"displayName\"\n",
		},
		{
			"with inner elements",
			filledElement,
			"dummy the_id <<stereo>> as \"displayName\" {\nDummy Child\n}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := NewWriter()
			if err := tt.element.Render(&writer); err != nil {
				t.Error("Render() got error but wanted none")
			}
			got := writer.String()
			if !reflect.DeepEqual(got, tt.want) {
				escapeErrorf(t, "Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewElementWithChildren(t *testing.T) {
	type args struct {
		typeString string
		id         string
		config     *ElementConfig
	}
	tests := []struct {
		name string
		args args
		want *ElementWithChildren
	}{
		{
			"valid id",
			args{
				typeString: "dummy",
				id:         "the_id",
				config:     nil,
			},
			&ElementWithChildren{
				Id:         "the_id",
				TypeString: "dummy",
				Config:     &ElementConfig{},
				Children:   []Renderable{},
			},
		},
		{
			"invalid id gets sanitized",
			args{
				typeString: "dummy",
				id:         "the:id",
				config:     nil,
			},
			&ElementWithChildren{
				Id:         "the_id",
				TypeString: "dummy",
				Config:     &ElementConfig{},
				Children:   []Renderable{},
			},
		},
		{
			"displayName and stereotype are set",
			args{
				typeString: "dummy",
				id:         "the_id",
				config:     &ElementConfig{DisplayName: "displayName", Stereotype: "stereo"},
			},
			&ElementWithChildren{
				TypeString: "dummy",
				Id:         "the_id",
				Config:     &ElementConfig{DisplayName: "displayName", Stereotype: "stereo"},
				Children:   []Renderable{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewElementWithChildren(tt.args.typeString, tt.args.id, tt.args.config)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewElementWithChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
