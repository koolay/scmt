package parse

import (
	"fmt"
	"io/ioutil"
	"strconv"

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
					//api := parseApi(comment)
					apiName := parseApiName(comment)
					apiVersion := parseApiVersion(comment)
					params := parseApiParam(comment)
					responses := parseResponse(comment)

					fmt.Println("---- response -----")
					for _, resp := range responses {
						fmt.Printf("code: %s, content: %s \n", strconv.Itoa(resp.Code), resp.Content)
					}
					fmt.Println("----response end -----")

					fmt.Printf("apiName: %s, apiVersion: %s \n", apiName, apiVersion)
					fmt.Println(params)
				}
			}
		}
	}
	return rtv
}
