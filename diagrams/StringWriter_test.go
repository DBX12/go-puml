package diagrams

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestNewWriter(t *testing.T) {
	actual := NewWriter()
	actualType := reflect.TypeOf(actual)
	wantedType := reflect.TypeOf(Writer{})
	if actualType != wantedType {
		t.Errorf("NewWriter() got type = %v, wantType %v", actualType, wantedType)
	}
	if actual.err != nil {
		t.Error("NewWriter() err was initialized with non-nil value")
	}
}

func TestWriter_GetError(t *testing.T) {
	expectedError := errors.New("error for testing")
	writer := Writer{
		buffer: bytes.Buffer{},
		err:    expectedError,
	}
	actualError := writer.GetError()
	if !reflect.DeepEqual(expectedError, actualError) {
		t.Errorf("GetError() = %v, want %v", actualError, expectedError)
	}
}

func TestWriter_Print(t *testing.T) {
	type fields struct {
		buffer bytes.Buffer
		err    error
	}
	type args struct {
		str string
	}
	filledBuffer := bytes.Buffer{}
	filledBuffer.Write([]byte("The first line\n"))
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"Empty buffer",
			fields{
				buffer: bytes.Buffer{},
				err:    nil,
			},
			args{"A single line"},
			"A single line",
		},
		{
			"Filled buffer",
			fields{
				buffer: filledBuffer,
				err:    nil,
			},
			args{"A single line"},
			"The first line\nA single line",
		},
		{
			"Do not write if an error exists",
			fields{
				buffer: bytes.Buffer{},
				err:    errors.New("an error"),
			},
			args{"A single line"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Writer{
				buffer: tt.fields.buffer,
				err:    tt.fields.err,
			}
			s.Print(tt.args.str)
			actual := s.String()
			if !reflect.DeepEqual(tt.want, actual) {
				t.Errorf("Print() = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestWriter_Printf(t *testing.T) {
	type fields struct {
		buffer bytes.Buffer
		err    error
	}
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"No formatting params",
			fields{
				buffer: bytes.Buffer{},
				err:    nil,
			},
			args{
				format: "Hello world",
				args:   nil,
			},
			"Hello world",
		},
		{
			"With formatting params",
			fields{
				buffer: bytes.Buffer{},
				err:    nil,
			},
			args{
				format: "Hello %s, your name has %d characters",
				args:   []interface{}{"John", 4},
			},
			"Hello John, your name has 4 characters",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Writer{
				buffer: tt.fields.buffer,
				err:    tt.fields.err,
			}
			s.Printf(tt.args.format, tt.args.args...)
			actual := s.String()
			if !reflect.DeepEqual(tt.want, actual) {
				t.Errorf("Printf() = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestWriter_Println(t *testing.T) {
	type fields struct {
		buffer bytes.Buffer
		err    error
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"With empty string",
			fields{
				buffer: bytes.Buffer{},
				err:    nil,
			},
			args{""},
			"\n",
		},
		{
			"With string",
			fields{
				buffer: bytes.Buffer{},
				err:    nil,
			},
			args{"A single line"},
			"A single line\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Writer{
				buffer: tt.fields.buffer,
				err:    tt.fields.err,
			}
			s.Println(tt.args.str)
			actual := s.String()
			if !reflect.DeepEqual(tt.want, actual) {
				t.Errorf("Println() = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestWriter_String(t *testing.T) {
	writer := Writer{
		buffer: bytes.Buffer{},
		err:    nil,
	}
	expectedString := "Testing content for the buffer"
	writer.buffer.Write([]byte(expectedString))
	actual := writer.String()
	if !reflect.DeepEqual(expectedString, actual) {
		t.Errorf("String() = %v, want %v", actual, expectedString)
	}
}
