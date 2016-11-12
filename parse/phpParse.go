package parse

import "github.com/go-openapi/spec"

type PhpParser struct {
}

func (*PhpParser) ParseComment(comment string) spec.PathItem {
	return spec.PathItem{}
}
