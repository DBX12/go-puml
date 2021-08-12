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
	_, err := s.buffer.Write([]byte(str))
	if err != nil {
		s.err = err
	}
}

func (s *Writer) Printf(format string, args ...interface{}) {
	if s.err != nil {
		return
	}
	_, err := s.buffer.Write([]byte(fmt.Sprintf(format, args...)))
	if err != nil {
		s.err = err
	}
}
