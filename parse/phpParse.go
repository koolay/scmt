package parse

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/spec"
)

type PhpParser struct {
}

var PHP_EXT = ".php"

func (parser *PhpParser) Parse(source string) map[string]spec.PathItem {

	files, err := findFiles(source, PHP_EXT)
	if err != nil {
		fmt.Printf("read files failed. source: %s", source)
		panic(err)
	}

	if len(files) < 1 {
		return nil
	}

	swaggerPathItems := map[string]spec.PathItem{}

	for _, f := range files {
		fmt.Printf("read file: %s \n", f)
		content, err := ioutil.ReadFile(f)
		if err == nil {
			sourceCode := string(content[:])
			comments := parseComments(sourceCode)
			for _, comment := range comments {
				api, err := parseApi(comment)
				if err != nil {
					fmt.Println("invalid api")
					continue
				}

				method := strings.ToUpper(api.Method)
				operate := spec.Operation{}
				apiName := parseApiName(comment)
				fmt.Println("---------process API: ", apiName, "---------")
				params := parseApiParam(comment)
				responses := parseResponse(comment)

				for _, param := range params {
					swaggerParam := spec.Parameter{}
					swaggerParam.Type = param.Type
					swaggerParam.Default = param.DefaultValue
					swaggerParam.Required = !param.Optional
					swaggerParam.Name = param.Field
					swaggerParam.Description = param.Description
					swaggerParam.Maximum = param.MaxNum
					swaggerParam.Minimum = param.MinNum
					swaggerParam.MaxLength = param.MaxLength
					swaggerParam.MinLength = param.MinLength
					operate.AddParam(&swaggerParam)

					/*
						Type         string
						Size         string
						Optional     bool
						Field        string
						DefaultValue string
						Description  string
					*/
				}

				swaggerResponses := spec.Responses{}
				swaggerResponses.StatusCodeResponses = map[int]spec.Response{}
				for _, resp := range responses {
					swaggerResp := spec.Response{}
					if resp.Content != "" {
						swaggerResponseSchema := &spec.Schema{}
						jsonSchemaBytes := ConvertResponseContentToJsonSchema(resp.Content)
						if err = swaggerResponseSchema.UnmarshalJSON(jsonSchemaBytes); err != nil {
							panic(err)
						}
						swaggerResp.Schema = swaggerResponseSchema
					} else {
						swaggerResp.Schema = nil
					}
					swaggerResponses.StatusCodeResponses[resp.Code] = swaggerResp
				}
				operate.Responses = &swaggerResponses
				swaggerPathItem := spec.PathItem{}
				switch method {
				case "GET":
					swaggerPathItem.Get = &operate
				case "POST":
					swaggerPathItem.Post = &operate
				case "PUT":
					swaggerPathItem.Put = &operate
				case "DELETE":
					swaggerPathItem.Delete = &operate
				case "PATCH":
					swaggerPathItem.Patch = &operate
				case "OPTIONS":
					swaggerPathItem.Options = &operate
				case "HEAD":
					swaggerPathItem.Head = &operate
				default:
					panic(fmt.Sprintf("not support method: %s", method))
				}

				// apiVersion := parseApiVersion(comment)
				swaggerPathItems[api.Path] = swaggerPathItem
			}
		}
	}
	return swaggerPathItems
}
