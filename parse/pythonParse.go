package parse

import "github.com/go-openapi/spec"

type PythonParser struct {
}

func (*PythonParser) Parse(source string) map[string]spec.PathItem {
	return map[string]spec.PathItem{}
}
