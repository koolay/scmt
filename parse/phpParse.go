package parse

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/spec"
)

type PhpParser struct {
}

var PHP_EXT = ".php"

func (parser *PhpParser) Parse(source string) []spec.PathItem {

	rtv := []spec.PathItem{}
	files, err := findFiles(source, PHP_EXT)
	if err == nil {

		for _, f := range files {
			fmt.Printf("read file: %s \n", f)
			content, err := ioutil.ReadFile(f)
			if err == nil {
				sourceCode := string(content[:])
				comments := parseComments(sourceCode)
				for _, comment := range comments {
					api := parseApi(comment)
					fmt.Println(api)
				}
			}
		}
	}
	return rtv
}
