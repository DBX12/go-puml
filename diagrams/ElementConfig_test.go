package diagrams

import (
	"reflect"
	"testing"
)

func TestElementConfig_render(t *testing.T) {
	tests := []struct {
		name   string
		config ElementConfig
		want   string
	}{
		{
			"no setting set",
			ElementConfig{},
			"",
		},
		{
			"single setting set",
			ElementConfig{
				BodyColor: "red",
			},
			" #red",
		},
		{
			"all settings set",
			ElementConfig{
				DisplayName: "DisplayName",
				Stereotype:  "Stereo",
				BodyColor:   "green",
				LineType:    "dotted",
				LineColor:   "blue",
				TextColor:   "white",
			},
			" #green;line:blue;line.dotted;text:white",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := NewWriter()
			tt.config.render(&writer)
			got := writer.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("render() string = %v, want = %v", got, tt.want)
			}
		})
	}
}
