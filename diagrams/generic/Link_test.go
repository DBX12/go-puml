package generic

import (
	"github.com/dbx12/go-puml/diagrams"
	"reflect"
	"testing"
)

func TestLink_Render(t *testing.T) {
	type fields struct {
		leftId  string
		rightId string
		config  *LinkConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"without label",
			fields{
				leftId:  "leftId",
				rightId: "rightId",
				config: &LinkConfig{
					Line:  "--",
					Label: "",
				},
			},
			"leftId -- rightId\n",
		},
		{
			"with label",
			fields{
				leftId:  "leftId",
				rightId: "rightId",
				config: &LinkConfig{
					Line:  "<-->",
					Label: "my label",
				},
			},
			"leftId <--> rightId : my label\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{
				leftId:  tt.fields.leftId,
				rightId: tt.fields.rightId,
				config:  tt.fields.config,
			}
			writer := diagrams.NewWriter()
			if err := l.Render(&writer); err != nil {
				t.Errorf("Render() got error and wanted none")
			}
			got := writer.String()
			if !reflect.DeepEqual(got, tt.want) {
				escapeErrorf(t, "Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLink(t *testing.T) {
	left := DummyLinkable{Id: "file_left"}
	right := DummyLinkable{Id: "file_right"}
	config := &LinkConfig{
		Line:  "<-->",
		Label: "my label",
	}
	type args struct {
		left   diagrams.Linkable
		right  diagrams.Linkable
		config *LinkConfig
	}
	tests := []struct {
		name string
		args args
		want Link
	}{
		{
			"without config uses default config",
			args{
				left:   left,
				right:  right,
				config: nil,
			},
			Link{
				leftId:  left.GetId(),
				rightId: right.GetId(),
				config: &LinkConfig{
					"--",
					"",
				},
			},
		},
		{
			"with config",
			args{
				left:   left,
				right:  right,
				config: config,
			},
			Link{
				leftId:  left.GetId(),
				rightId: right.GetId(),
				config:  config,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := *NewLink(tt.args.left, tt.args.right, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLink() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.config, tt.want.config) {
				t.Errorf("NewLink() config = %v, want %v", got.config, tt.want.config)
			}
			if !reflect.DeepEqual(got.leftId, tt.want.leftId) {
				t.Errorf("NewLink() leftId = %v, want %v", got.leftId, tt.want.leftId)
			}
			if !reflect.DeepEqual(got.rightId, tt.want.rightId) {
				t.Errorf("NewLink() rightId = %v, want %v", got.rightId, tt.want.rightId)
			}
		})
	}
}

func TestNewLinkFromIds(t *testing.T) {
	config := &LinkConfig{
		Line:  "<-->",
		Label: "my label",
	}
	type args struct {
		leftId  string
		rightId string
		config  *LinkConfig
	}
	tests := []struct {
		name string
		args args
		want Link
	}{
		{
			"without config uses default config",
			args{
				leftId:  "leftId",
				rightId: "rightId",
				config:  nil,
			},
			Link{
				leftId:  "leftId",
				rightId: "rightId",
				config: &LinkConfig{
					Line:  "--",
					Label: "",
				},
			},
		},
		{
			"with config",
			args{
				leftId:  "leftId",
				rightId: "rightId",
				config:  config,
			},
			Link{
				leftId:  "leftId",
				rightId: "rightId",
				config:  config,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := *NewLinkFromIds(tt.args.leftId, tt.args.rightId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkFromIds() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.config, tt.want.config) {
				t.Errorf("NewLinkFromIds() config = %v, want %v", got.config, tt.want.config)
			}
			if !reflect.DeepEqual(got.leftId, tt.want.leftId) {
				t.Errorf("NewLinkFromIds() leftId = %v, want %v", got.leftId, tt.want.leftId)
			}
			if !reflect.DeepEqual(got.rightId, tt.want.rightId) {
				t.Errorf("NewLinkFromIds() rightId = %v, want %v", got.rightId, tt.want.rightId)
			}
		})
	}
}
