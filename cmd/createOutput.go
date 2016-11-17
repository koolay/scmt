package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/spec"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

type OutPuter struct {
	Swagger          *spec.Swagger
	OutputFlags      []string
	swaggerJsonBytes []byte
}

func (out *OutPuter) Output() error {

	if len(out.OutputFlags) < 1 {
		out.OutputFlags = []string{OUTPUT_STDOUT}
	}

	out.OutputFlags = removeDuplicatesUnordered(out.OutputFlags)
	serialize(out)
	var err error
	for _, flag := range out.OutputFlags {
		dest := parseOutputFlag(flag)

		switch dest {
		case OUTPUT_API:
			err = out.toApi(flag)
		case OUTPUT_YML:
			err = out.toYml(flag)
		case OUTPUT_JSON:
			err = out.toJson(flag)
		case OUTPUT_STDOUT:
			err = out.toStdout()
		default:
			err = out.toStdout()
		}
	}

	return err

}

func serialize(out *OutPuter) {

	bytes, err := out.Swagger.MarshalJSON()
	if err != nil {
		panic(err)
	}
	out.swaggerJsonBytes = bytes
}

func (out *OutPuter) toYml(dest string) error {

	fmt.Println("-------------output to yml-------------")
	return nil
}

func (out *OutPuter) toJson(dest string) error {
	fmt.Println("-------------output to json-------------")
	fullname := dest
	dir, filename := filepath.Split(fullname)
	if dir == "" {
		tmpDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return err
		}
		dir = tmpDir
		fullname = fmt.Sprintf("%s/%s", dir, filename)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return errors.New(1, fmt.Sprintf("%s not exists", dir))
	}

	if file, err := os.Create(fullname); err == nil {
		defer file.Close()
		file.Write(out.swaggerJsonBytes)
		fmt.Println("file saved: ", fullname)
	} else {
		return err
	}

	return nil

}

func (out *OutPuter) toApi(dest string) error {
	fmt.Println("-------------output to api-------------")
	payload := map[string]interface{}{
		"swagger": string(out.swaggerJsonBytes),
	}
	headers := viper.Get("headers").([]string)
	request := gorequest.New()
	request.Put(dest).SendMap(payload)
	for _, v := range headers {
		val := strings.Replace(v, `"`, "", -1)
		pies := strings.Split(val, "=")
		if len(pies) == 2 {
			headerKey := strings.TrimSpace(pies[0])
			headerVal := strings.TrimSpace(pies[1])
			request.Set(headerKey, headerVal)
		} else {
			return errors.New(1, "invalid args of -H")
		}
	}
	resp, body, errs := request.End()
	if errs != nil {
		return errs[0]
	}
	fmt.Println("http status:", resp.StatusCode)
	fmt.Println(body)
	return nil
}

func (out *OutPuter) toStdout() error {
	fmt.Println("-------------output to stdout-------------")
	fmt.Println(string(out.swaggerJsonBytes))
	return nil
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
