package generic

import (
	"github.com/dbx12/go-puml/diagrams"
	"reflect"
	"testing"
)

func TestComment_Render(t *testing.T) {
	type fields struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Single line comment",
			fields{text: "Hello comment"},
			"' Hello comment\n",
		},
		{
			"Multi line comment",
			fields{text: "Hello comment\nOn two lines"},
			"' Hello comment\n' On two lines\n",
		},
		{
			"With trailing newline",
			fields{text: "Hello comment\n\n"},
			"' Hello comment\n",
		},
		{
			"Comment with embedded empty line",
			fields{text: "Hello comment\n\nThird line"},
			"' Hello comment\n' \n' Third line\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Comment{
				text: tt.fields.text,
			}
			writer := diagrams.NewWriter()
			if err := c.Render(&writer); err != nil {
				t.Error("Render() got error but wanted none")
			}
			actual := writer.String()
			if !reflect.DeepEqual(tt.want, actual) {
				escapeErrorf(t, "Render() string = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestNewComment(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want *Comment
	}{
		{
			"Initialize comment object",
			args{text: "Hello comment"},
			&Comment{text: "Hello comment"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComment(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
