package diagrams

import (
	"reflect"
	"strings"
	"testing"
)

// This file contains helper functions used during testing

func escapeErrorf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	t.Errorf(format, escapeMessage(t, args)...)
}

func escapeFatalf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	t.Fatalf(format, escapeMessage(t, args)...)
}

func escapeMessage(t *testing.T, args []interface{}) []interface{} {
	t.Helper()
	var escaped []interface{}
	for _, arg := range args {
		if reflect.TypeOf(arg).String() == "string" {
			escaped = append(escaped, strings.ReplaceAll(arg.(string), "\n", "\\n"))
		} else {
			escaped = append(escaped, arg)
		}
	}
	return escaped
}

type DummyRenderable struct {
	Content string
}

func (d DummyRenderable) Render(writer *Writer) error {
	writer.Println(d.Content)
	return writer.GetError()
}

type DummyLinkable struct {
	DummyRenderable
	Id string
}

func (d DummyLinkable) GetId() string {
	return d.Id
}

type DummyContainer struct {
	DummyRenderable
	children []Renderable
}

func (d DummyContainer) Add(r ...Renderable) {
	d.children = append(d.children, r...)
}
