package parse

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"
)

var reCommentBlock = regexp.MustCompile(`/\*(.|[\r\n])*?\*/`)
var reApi = regexp.MustCompile(`@api\s\{(?P<method>\w+)\}\s+(?P<path>(/[^\s]+)+)(\s(?P<title>[^\*\r\n]+))?`)
var reApiName = regexp.MustCompile(`@apiName\s+(?P<name>[^\r\n]+)`)
var reApiVersion = regexp.MustCompile(`@apiVersion\s+(?P<name>[^\r\n]+)`)
var reApiParam = regexp.MustCompile(`@apiParam\s+(?P<name>[^\r\n]+)`)
var reApiParamGroup = regexp.MustCompile(`(?P<type>\{.+\})\s+(?P<field>(\[[^\s=\[]+\=?[^\=\]]+\])|([^\s=\[]+\=?[^\s\=\]]+))(\s+(?P<descriptionoptional>[^\r\n]*))?`)
var reApiParamTypeGroup = regexp.MustCompile(`\{\s*(?P<type>[a-z]+)(\{(?P<range>(\d+)?(-|\.+)?\d+)\})?\s*`)
var reApiParamFieldGroup = regexp.MustCompile(`\[?\s*(?P<paramName>[^=\s\[\]]+)(=\s*"?(?P<default>[^="\[\]]+)"?)?`)
var reApiResponse = regexp.MustCompile(`(?ms)@apiResponse\s+(?P<responseCode>\d+)\s*(?P<content>\{.*?\*\s*\})`)

type Parser interface {
	Parse(source string) []spec.PathItem
}

type Api struct {
	Method string
	Path   string
	Title  string
}

type Param struct {
	Type         string
	Size         string
	Optional     bool
	Field        string
	DefaultValue string
	Description  string
}

type Response struct {
	Code    int
	Content string
}

func parseApi(sourceCode string) Api {
	result := matchGroup(reApi, sourceCode)
	return Api{Path: result["path"], Title: result["title"], Method: result["method"]}
}

func parseApiVersion(sourceCode string) string {
	match := reApiVersion.FindStringSubmatch(sourceCode)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

func parseApiName(sourceCode string) string {
	match := reApiName.FindStringSubmatch(sourceCode)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

func parseApiParam(commentBlock string) []*Param {

	rtv := []*Param{}
	// get whole param string
	match := reApiParam.FindAllStringSubmatch(commentBlock, -1)
	params := []string{}
	for _, v := range match {
		if len(v) > 1 {
			params = append(params, v[1])
		}
	}

	for _, v := range params {

		fmt.Printf("param: %s \n", v)
		match := matchGroup(reApiParamGroup, v)
		if len(match) > 0 {
			param := new(Param)
			typeString := match["type"]   // {integer{100-200}}  or {string{..5}}
			fieldString := match["field"] // [name=abc] or id=123
			param.Description = match["description"]
			optional := false
			var paramName string
			var defaultValue string

			if fieldString == "" {
				log.Fatal("parse field failed! source: \n %s \n", commentBlock)
			}

			if typeString == "" {
				log.Fatal("parse failed! invalid type, source: \n %s \n", commentBlock)
			}

			matchFieldItems := matchGroup(reApiParamFieldGroup, fieldString)
			paramName = matchFieldItems["paramName"]
			defaultValue = matchFieldItems["default"]
			if strings.HasPrefix(fieldString, "[") {
				optional = true
			}

			matchTypeItems := matchGroup(reApiParamTypeGroup, typeString)
			paramType := "string"
			paramRange := "" // 1-2 or -2 or 1- or 1.. or 1..2 or ..2
			if len(matchTypeItems) > 0 {
				paramType = matchTypeItems["type"]
				paramRange = matchTypeItems["range"]
			}
			param.DefaultValue = defaultValue
			param.Field = paramName
			param.Optional = optional
			param.Size = paramRange
			param.Type = paramType
			rtv = append(rtv, param)
		}
	}

	return rtv

}

func parseResponse(commentBlock string) []Response {

	result := reApiResponse.FindAllStringSubmatch(commentBlock, -1)
	rtv := []Response{}
	starRe := regexp.MustCompile(`\s*\*\s*`)
	for _, v := range result {
		response := make(map[string]string)
		for i, name := range reApiResponse.SubexpNames() {
			if i != 0 {
				response[name] = v[i]
			}
		}
		resp := Response{}
		responseCode := response["responseCode"]
		if code, err := strconv.Atoi(responseCode); err == nil {
			resp.Code = code
		} else {
			log.Fatal("invalid response code %s", response["responseCode"])
		}
		resp.Content = starRe.ReplaceAllString(response["content"], "")
		rtv = append(rtv, resp)
	}
	return rtv
}

func matchGroup(re *regexp.Regexp, text string) map[string]string {
	match := re.FindStringSubmatch(text)
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}
	return result
}

func recursiveFindFiles(root string, pattern string) ([]string, error) {
	files := make([]string, 0)
	findfile := func(path string, f os.FileInfo, err error) (inner error) {
		if err != nil {
			return
		}
		if f.IsDir() {
			return
		} else if match, innerr := filepath.Match(pattern, f.Name()); innerr == nil && match {
			files = append(files, path)
		}
		return
	}
	err := filepath.Walk(root, findfile)
	if len(files) == 0 {
		return files, err
	} else {
		return files, err
	}
}

// find files from folder
func findFiles(source string, ext string) ([]string, error) {

	files := []string{}
	if ext == "" {
		return files, nil
	}

	pattern := fmt.Sprintf("*%s", ext)
	info, err := os.Stat(source)
	if os.IsNotExist(err) {
		log.Println(err.Error())
		return files, err
	} else if info.IsDir() {
		fmt.Printf("start search dir: %s, match: %s \n", source, pattern)
		foundFiles, err := recursiveFindFiles(source, pattern)
		if err == nil {
			files = append(files, foundFiles...)
		}
	} else {
		files = append(files, source)
	}
	return files, err

}

func parseComments(sourceCode string) []string {
	return reCommentBlock.FindAllString(sourceCode, -1)
}
