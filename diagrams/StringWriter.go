package diagrams

import (
	"bytes"
	"fmt"
)

type Writer struct {
	buffer bytes.Buffer
	err    error
}

func NewWriter() Writer {
	return Writer{
		buffer: bytes.Buffer{},
		err:    nil,
	}
}

func (s *Writer) String() string {
	return s.buffer.String()
}

func (s *Writer) GetError() error {
	return s.err
}

func (s *Writer) Println(str string) {
	s.Print(str + "\n")
}

func (s *Writer) Print(str string) {
	if s.err != nil {
		return
	}
	// no need to check for an error the Write() method of bytes.Buffer will
	// never return an error but panic directly
	s.buffer.Write([]byte(str))
}

func (s *Writer) Printf(format string, args ...interface{}) {
	s.Print(fmt.Sprintf(format, args...))
}
