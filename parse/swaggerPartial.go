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
	"github.com/spf13/viper"
)

// var reCommentBlock = regexp.MustCompile(`/\*(.|[\r\n])*?\*/`)
var reApi = regexp.MustCompile(`@api\s\{(?P<method>\w+)\}\s+(?P<path>(/[^\s]+)+)(\s(?P<title>[^\*\r\n]+))?`)
var reApiModule = regexp.MustCompile(`@apiGroup\s+(?P<module>[^\r\n\s]+)`)

//var reApiName = regexp.MustCompile(`@apiName\s+(?P<name>[^\r\n]+)`)
var reApiVersion = regexp.MustCompile(`@apiVersion\s+(?P<name>[^\r\n]+)`)
var reApiParam = regexp.MustCompile(`@apiParam\s+(?P<name>[^\r\n]+)`)
var reApiParamGroup = regexp.MustCompile(`(?P<in>[a-zA-Z]+)\s+(?P<type>\{.+\})\s+(?P<field>(\[[^\s=\[]+\=?[^\=\]]+\])|([^\s=\[]+\=?[^\s\=\]]+))(\s+(?P<description>[^\r\n]*))?`)
var reApiParamTypeGroup = regexp.MustCompile(`\{\s*(?P<type>[a-z]+)(\{\s*(((?P<min>\d+)?\-(?P<max>\d+)?)|((?P<short>\d+)?\.\.(?P<length>\d+)?))\s*\})?\s*\}`)
var reApiParamFieldGroup = regexp.MustCompile(`\[?\s*(?P<paramName>[^=\s\[\]]+)(=\s*"?(?P<default>[^="\[\]]+)"?)?`)
var reApiResponse = regexp.MustCompile(`(?ms)@apiResponse\s+(?P<responseCode>\d+)(\s*(?P<content>(\{.*?\*\s*\})|(\[.*?\*\s*\])))?`)

type Parser struct {
	lang                string
	sourceCode          string
	commentRegexPattern string

	Swagger *spec.Swagger
}

var commentRegexPatternMap = map[string]string{"javascript": `/\*(.|[\r\n])*?\*/`, "php": `/\*(.|[\r\n])*?\*/`, "python": `\"\"\"(.|[\r\n])*?\"\"\"`}

func NewParser(lang string) (parser *Parser) {
	parser = &Parser{lang: lang, commentRegexPattern: commentRegexPatternMap[lang]}

	// init swagger
	parser.Swagger = &spec.Swagger{}
	swaggerInfo := spec.Info{}
	swaggerInfo.Title = viper.Get("name").(string)
	swaggerInfo.Version = viper.Get("version").(string)
	parser.Swagger.Info = &swaggerInfo
	parser.Swagger.BasePath = "/"
	parser.Swagger.Swagger = "2.0"
	parser.Swagger.Definitions = spec.Definitions{}
	parser.Swagger.Paths = &spec.Paths{}
	parser.Swagger.Paths.Paths = map[string]spec.PathItem{}

	return parser
}

type Api struct {
	Method string
	Path   string
	Title  string
}

type Param struct {
	In           string
	Type         string
	Optional     bool
	Field        string
	DefaultValue interface{}
	Description  string
	MaxNum       *float64
	MinNum       *float64
	MaxLength    *int64
	MinLength    *int64
}

type Response struct {
	Code        int
	Content     string
	Description string
}

// return module or tag
func (parser *Parser) parseApiTag(commentBlock string) string {

	match := reApiModule.FindStringSubmatch(commentBlock)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

func (parser *Parser) parseApi(sourceCode string) (Api, error) {
	result := matchGroup(reApi, sourceCode)
	if len(result) == 0 {
		return Api{}, &ParseError{}
	}
	return Api{Path: result["path"], Title: result["title"], Method: result["method"]}, nil
}

func (parser *Parser) parseApiVersion(sourceCode string) string {
	match := reApiVersion.FindStringSubmatch(sourceCode)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

func (parser *Parser) parseApiParam(commentBlock string) []*Param {

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

		match := matchGroup(reApiParamGroup, v)
		if len(match) > 0 {
			param := new(Param)
			in := match["in"]
			typeString := match["type"]   // {integer{100-200}}  or {string{..5}}
			fieldString := match["field"] // [name=abc] or id=123
			optional := false
			var paramName string
			var defaultValue interface{}

			if fieldString == "" {
				log.Fatal("parse field failed! source: \n %s \n", commentBlock)
			}

			if typeString == "" {
				log.Fatal("parse failed! invalid type, source: \n %s \n", commentBlock)
			}

			matchFieldItems := matchGroup(reApiParamFieldGroup, fieldString)
			paramName = matchFieldItems["paramName"]
			if matchFieldItems["default"] == "" {
				defaultValue = nil
			} else {
				defaultValue = matchFieldItems["default"]
			}
			if strings.HasPrefix(fieldString, "[") {
				optional = true
			}

			matchTypeItems := matchGroup(reApiParamTypeGroup, typeString)
			paramType := "string"
			if len(matchTypeItems) > 0 {
				paramType = matchTypeItems["type"]
				if matchTypeItems["max"] != "" {
					maxNum, _ := strconv.ParseFloat(matchTypeItems["max"], 64)
					param.MaxNum = &maxNum
				}
				if matchTypeItems["min"] != "" {
					minNum, _ := strconv.ParseFloat(matchTypeItems["min"], 64)
					param.MinNum = &minNum
				}
				if matchTypeItems["short"] != "" {
					minLength, _ := strconv.ParseInt(matchTypeItems["short"], 10, 64)
					param.MinLength = &minLength
				}
				if matchTypeItems["length"] != "" {
					maxLength, _ := strconv.ParseInt(matchTypeItems["length"], 10, 64)
					param.MaxLength = &maxLength
				}
			}

			if match["description"] != "" {
				param.Description = match["description"]
			} else {
				param.Description = paramName
			}
			param.In = in
			param.DefaultValue = defaultValue
			param.Field = paramName
			param.Optional = optional
			param.Type = paramType
			rtv = append(rtv, param)
		}
	}

	return rtv

}

func (parser *Parser) parseResponse(commentBlock string) []Response {

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
		resp.Description = " "
		rtv = append(rtv, resp)
	}
	return rtv
}

func matchGroup(re *regexp.Regexp, text string) map[string]string {
	result := make(map[string]string)
	match := re.FindStringSubmatch(text)
	if len(match) < 2 {
		return result
	}
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
func FindFiles(source string, ext string) ([]string, error) {

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

func (parser *Parser) parseComments(sourceCode string) []string {
	var reCommentBlock = regexp.MustCompile(parser.commentRegexPattern)
	return reCommentBlock.FindAllString(sourceCode, -1)
}
