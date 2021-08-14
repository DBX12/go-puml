package diagrams

import (
	"reflect"
	"testing"
)

func TestNewNoteForId(t *testing.T) {
	type args struct {
		id       string
		position NotePosition
		contents []string
	}
	tests := []struct {
		name string
		args args
		want *Note
	}{
		{
			"Successful",
			args{
				id:       "targetElement",
				position: ABOVE,
				contents: []string{"First note line", "Second line"},
			},
			&Note{
				linkedTo: "targetElement",
				position: ABOVE,
				contents: []string{"First note line", "Second line"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoteForId(tt.args.id, tt.args.position, tt.args.contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteForId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNoteForLinkable(t *testing.T) {
	type args struct {
		r        Linkable
		position NotePosition
		contents []string
	}
	tests := []struct {
		name string
		args args
		want *Note
	}{
		{
			"Successful",
			args{
				r: DummyLinkable{
					Id: "dummy_id",
				},
				position: ABOVE,
				contents: []string{"First note line", "Second line"},
			},
			&Note{
				linkedTo: "dummy_id",
				position: ABOVE,
				contents: []string{"First note line", "Second line"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoteForLinkable(tt.args.r, tt.args.position, tt.args.contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteForLinkable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNote_Render(t *testing.T) {
	type fields struct {
		linkedTo string
		position NotePosition
		contents []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Successful",
			fields{
				linkedTo: "targetId",
				position: ABOVE,
				contents: []string{"Line 1", "----", "Line 2"},
			},
			"note top of targetId\nLine 1\n----\nLine 2\nendnote\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Note{
				linkedTo: tt.fields.linkedTo,
				position: tt.fields.position,
				contents: tt.fields.contents,
			}
			writer := NewWriter()
			if err := n.Render(&writer); err != nil {
				t.Error("Render() got error but wanted none")
			}
			got := writer.String()
			if !reflect.DeepEqual(got, tt.want) {
				escapeErrorf(t, "Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
