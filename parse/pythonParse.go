package parse

import "github.com/go-openapi/spec"

type PythonParser struct {
}

func (*PythonParser) ParseComment(comment string) spec.PathItem {
	return spec.PathItem{}

}
