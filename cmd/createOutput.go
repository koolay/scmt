package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-openapi/spec"
)

type OutPuter struct {
	Swagger          *spec.Swagger
	OutputFlags      []string
	swaggerJsonBytes []byte
}

func (out *OutPuter) Output() {

	if len(out.OutputFlags) < 1 {
		out.OutputFlags = []string{OUTPUT_STDOUT}
	}

	out.OutputFlags = removeDuplicatesUnordered(out.OutputFlags)
	serialize(out)

	for _, flag := range out.OutputFlags {
		dest := parseOutputFlag(flag)

		switch dest {
		case OUTPUT_API:
			out.toApi()
		case OUTPUT_YML:
			out.toYml()
		case OUTPUT_JSON:
			out.toJson()
		case OUTPUT_STDOUT:
			out.toStdout()
		default:
			out.toStdout()
		}
	}

}

func serialize(out *OutPuter) {

	bytes, err := out.Swagger.MarshalJSON()
	if err != nil {
		panic(err)
	}
	out.swaggerJsonBytes = bytes
}

func (out *OutPuter) toYml() {

	fmt.Println("-------------output to yml-------------")
}

func (out *OutPuter) toJson() {
	fmt.Println("-------------output to json-------------")
}

func (out *OutPuter) toApi() {
	fmt.Println("-------------output to api-------------")
}

func (out *OutPuter) toStdout() {
	fmt.Println("-------------output to stdout-------------")
	fmt.Println(string(out.swaggerJsonBytes))
}

func parseOutputFlag(cmdFlag string) string {

	reApiFlag := regexp.MustCompile(`^https?\/\/[^\s]+`)
	if reApiFlag.Match([]byte(cmdFlag)) {
		return OUTPUT_API
	} else if strings.HasSuffix(cmdFlag, ".json") {
		return OUTPUT_JSON
	} else if strings.HasSuffix(cmdFlag, ".yml") {
		return OUTPUT_YML
	} else {
		return OUTPUT_STDOUT
	}

}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}
