package diagrams

import (
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
