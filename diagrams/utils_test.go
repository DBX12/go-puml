package diagrams

import (
	"reflect"
	"testing"
)

func TestSanitizeId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"replace invalid chars",
			args{"a very-bad:identifier/really-bad"},
			"avery_bad_identifier_really_bad",
		},
		{
			"change nothing in valid id",
			args{"a_valid_identifier"},
			"a_valid_identifier",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeId(tt.args.id); got != tt.want {
				t.Errorf("SanitizeId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assertValidId(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Prevent whitespaces",
			args{"model foo bar"},
			true,
		},
		{
			"Allow good names",
			args{"model_foo"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := assertValidId(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("assertValidId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_conditionalPrintf(t *testing.T) {
	type args struct {
		format string
		value  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"With content",
			args{
				format: "-> %s",
				value:  "content",
			},
			"-> content",
		},
		{
			"Without content",
			args{
				format: "-> %s",
				value:  "",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := NewWriter()
			conditionalPrintf(&writer, tt.args.format, tt.args.value)
			got := writer.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conditionalPrintf() string = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_renderInnerElements(t *testing.T) {
	tests := []struct {
		name        string
		renderables []Renderable
		want        string
	}{
		{
			"without elements",
			[]Renderable{},
			"",
		},
		{
			"with elements",
			[]Renderable{
				DummyRenderable{Content: "Dummy 01"},
				DummyRenderable{Content: "Dummy 02"},
			},
			" {\nDummy 01\nDummy 02\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := NewWriter()
			if err := renderInnerElements(&writer, tt.renderables); err != nil {
				t.Error("renderInnerElements() got an error but wanted none")
			}
			got := writer.String()
			if !reflect.DeepEqual(got, tt.want) {
				escapeErrorf(t, "renderInnerElements() string = %v, want = %v", got, tt.want)
			}
		})
	}
}
