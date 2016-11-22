package parse

import (
	"fmt"
	"strings"

	"github.com/go-openapi/spec"
)

type PhpParser struct {
}

var PHP_EXT = ".php"

func (parser *Parser) Parse(sourceCode string) {

	//pathsMap := parser.Swagger.Paths.Paths
	comments := parser.parseComments(sourceCode)
	for _, comment := range comments {
		api, err := parser.parseApi(comment)
		if err != nil {
			fmt.Println("invalid api")
			continue
		}
		if api.Path == "" {
			panic(fmt.Sprintf("parse failed, path can not empty. code: \n %s \n", comment))
		}

		method := strings.ToUpper(api.Method)
		apiTitle := strings.TrimSpace(api.Title)
		operate := spec.Operation{}
		operate.Description = apiTitle
		apiTag := parser.parseApiTag(comment)
		params := parser.parseApiParam(comment)
		responses := parser.parseResponse(comment)

		fmt.Println("---------process API: ", apiTitle, "---------")

		for _, param := range params {
			swaggerParam := spec.Parameter{}
			swaggerParam.In = param.In
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
			swaggerResp.Description = resp.Description
			swaggerResponses.StatusCodeResponses[resp.Code] = swaggerResp
		}
		operate.Responses = &swaggerResponses
		operate.Tags = []string{apiTag}
		setOperation := func(item *spec.PathItem) {

			switch method {
			case "GET":
				item.Get = &operate
			case "POST":
				item.Post = &operate
			case "PUT":
				item.Put = &operate
			case "OPTIONS":
				item.Options = &operate
			case "DELETE":
				item.Delete = &operate
			case "PATCH":
				item.Patch = &operate
			case "HEAD":
				item.Head = &operate
			default:
				panic("not support method")
			}
		}

		pathItem, ok := parser.Swagger.Paths.Paths[api.Path]
		if ok {
			setOperation(&pathItem)
			parser.Swagger.Paths.Paths[api.Path] = pathItem
		} else {
			pathItem := spec.PathItem{}
			setOperation(&pathItem)
			parser.Swagger.Paths.Paths[api.Path] = pathItem
		}

		// apiVersion := parseApiVersion(comment)
	}
}
