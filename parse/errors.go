package parse

import "fmt"

type ParseError struct {
	Pattern  string
	Filename string
	Source   string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Parse %s failed, with: %s. filename: %s", e.Source, e.Pattern, e.Filename)
}
