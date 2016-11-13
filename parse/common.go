package parse

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-openapi/spec"
)

var reCommentBlock = regexp.MustCompile(`/\*(.|[\r\n])*?\*/`)
var reApi = regexp.MustCompile(`@api\s\{(?P<method>\w+)\}\s+(?P<path>(/[^\s]+)+)(\s(?P<title>[^\*\r\n]+))?`)

type Parser interface {
	Parse(source string) []spec.PathItem
}

type Api struct {
	Method string
	Path   string
	Title  string
}

type Param struct {
	group         string
	Type          string
	Size          string
	AllowedValues []string
	Optional      bool
	Field         string
	DefaultValue  string
	Description   string
}

func parseApi(sourceCode string) Api {

	result := matchGroup(reApi, sourceCode)
	return Api{Path: result["path"], Title: result["title"], Method: result["method"]}
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
