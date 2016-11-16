package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
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
			out.toApi(flag)
		case OUTPUT_YML:
			out.toYml(flag)
		case OUTPUT_JSON:
			out.toJson(flag)
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

func (out *OutPuter) toYml(dest string) {

	fmt.Println("-------------output to yml-------------")
}

func (out *OutPuter) toJson(dest string) {
	fmt.Println("-------------output to json-------------")
}

func (out *OutPuter) toApi(dest string) {
	fmt.Println("-------------output to api-------------")
	headers := viper.Get("headers").([]string)
	request := gorequest.New()
	for _, v := range headers {
		val := strings.Replace(v, `"`, "", -1)
		pies := strings.Split(val, "=")
		if len(pies) == 2 {
			headerKey := strings.TrimSpace(pies[0])
			headerVal := strings.TrimSpace(pies[1])
			request.Set(headerKey, headerVal)
			fmt.Printf("-H %s: %s \n", headerKey, headerVal)
		} else {
			panic("invalid args of -H")
		}
	}
	resp, body, errs := request.Put(dest).Send(`{"swagger": "` + string(out.swaggerJsonBytes) + `"}`).End()
	if errs != nil {
		panic(errs[0])
	}
	fmt.Println("http status:", resp.StatusCode)
	fmt.Println(body)
}

func (out *OutPuter) toStdout() {
	fmt.Println("-------------output to stdout-------------")
	fmt.Println(string(out.swaggerJsonBytes))
}

func parseOutputFlag(cmdFlag string) string {

	reApiFlag := regexp.MustCompile(`^https?:\/\/[^\s]+`)
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
