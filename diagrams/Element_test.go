package diagrams

import (
	"reflect"
	"testing"
)

func TestElement_GetId(t *testing.T) {
	want := "the_id"
	element := NewElement("dummy", want, nil)
	got := element.GetId()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetId() = %v, want %v", got, want)
	}
}

func TestElement_Render(t *testing.T) {
	tests := []struct {
		name    string
		element *Element
		want    string
	}{
		{
			"without displayname",
			NewElement("dummy", "the_id", nil),
			"dummy the_id\n",
		},
		{
			"with display name",
			NewElement("dummy", "the_id", &ElementConfig{DisplayName: "Line 1\nLine 2"}),
			"dummy the_id as \"Line 1\nLine 2\"\n",
		},
		{
			"with stereotype",
			NewElement("dummy", "the_id", &ElementConfig{DisplayName: "displayName", Stereotype: "stereo"}),
			"dummy the_id <<stereo>> as \"displayName\"\n",
		},
		{
			"with full config",
			NewElement("dummy", "the_id", &ElementConfig{
				DisplayName: "displayName",
				Stereotype:  "stereo",
				BodyColor:   "orange",
				LineType:    "dotted",
				LineColor:   "blue",
				TextColor:   "black",
			}),
			"dummy the_id <<stereo>> as \"displayName\" #orange;line:blue;line.dotted;text:black\n",
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

func TestNewElement(t *testing.T) {
	type args struct {
		typeString string
		id         string
		config     *ElementConfig
	}
	tests := []struct {
		name string
		args args
		want *Element
	}{
		{
			"valid id",
			args{
				typeString: "dummy",
				id:         "the_id",
				config:     nil,
			},
			&Element{
				TypeString: "dummy",
				Id:         "the_id",
				Config:     &ElementConfig{},
			},
		},
		{
			"invalid id gets sanitized",
			args{
				typeString: "dummy",
				id:         "the:id",
				config:     nil,
			},
			&Element{
				TypeString: "dummy",
				Id:         "the_id",
				Config:     &ElementConfig{},
			},
		},
		{
			"config is set",
			args{
				typeString: "dummy",
				id:         "id",
				config: &ElementConfig{
					DisplayName: "The displayName",
					Stereotype:  "stereo",
					BodyColor:   "green",
					LineType:    "dotted",
					LineColor:   "red",
					TextColor:   "white",
				},
			},
			&Element{
				TypeString: "dummy",
				Id:         "id",
				Config: &ElementConfig{
					DisplayName: "The displayName",
					Stereotype:  "stereo",
					BodyColor:   "green",
					LineType:    "dotted",
					LineColor:   "red",
					TextColor:   "white",
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewElement(tt.args.typeString, tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
