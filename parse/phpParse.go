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
					api, err := parseApi(comment)
					if err != nil {
						fmt.Println("invalid api")
						continue
					}

					method := strings.ToUpper(api.Method)
					operate := spec.Operation{}
					apiName := parseApiName(comment)
					apiVersion := parseApiVersion(comment)
					params := parseApiParam(comment)
					responses := parseResponse(comment)

					fmt.Println("---- process params -----")

					for _, param := range params {
						swaggerParam := spec.Parameter{}
						swaggerParam.Type = param.Type
						swaggerParam.Default = param.DefaultValue
						swaggerParam.Required = !param.Optional
						swaggerParam.Name = param.Field
						swaggerParam.Description = param.Description
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

					fmt.Println("---- response -----")
					swaggerResponses := spec.Responses{}
					for _, resp := range responses {
						swaggerResp := spec.Response{}
						jsonSchemaBytes := ConvertResponseContentToJsonSchema(resp.Content)
						swaggerResponseSchema := &spec.Schema{}
						if err = swaggerResponseSchema.UnmarshalJSON(jsonSchemaBytes); err != nil {
							panic(err)
						}
						swaggerResp.Schema = swaggerResponseSchema
						swaggerResponses.StatusCodeResponses = map[int]spec.Response{}
						swaggerResponses.StatusCodeResponses[resp.Code] = swaggerResp
					}
					operate.Responses = &swaggerResponses
					swaggerPath := spec.PathItem{}
					switch method {
					case "GET":
						swaggerPath.Get = &operate
					case "POST":
						swaggerPath.Post = &operate
					case "PUT":
						swaggerPath.Put = &operate
					case "DELETE":
						swaggerPath.Delete = &operate
					case "PATCH":
						swaggerPath.Patch = &operate
					case "OPTIONS":
						swaggerPath.Options = &operate
					case "HEAD":
						swaggerPath.Head = &operate
					default:
						panic(fmt.Sprintf("not support method: %s", method))
					}

					if swaggerPathJsonBytes, err := swaggerPath.MarshalJSON(); err == nil {
						fmt.Println(string(swaggerPathJsonBytes))
					} else {
						panic(err)
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
