package parse

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/go-openapi/spec"
)

type Parser interface {
	ParseComment(comment string) spec.PathItem
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

func ReadFiles(source string, ext string) (map[string]string, error) {
	rtv := make(map[string]string)
	if ext == "" {
		return rtv, nil
	}

	files := []string{}

	pattern := fmt.Sprintf("*%s", ext)
	info, err := os.Stat(source)
	if os.IsNotExist(err) {
		log.Println(err.Error())
		return rtv, err
	} else if info.IsDir() {
		fmt.Printf("start search dir: %s, match: %s \n", source, pattern)
		foundFiles, err := recursiveFindFiles(source, pattern)
		if err == nil {
			files = append(files, foundFiles...)
		} else {
			return rtv, err
		}
	} else {
		files = append(files, source)
	}

	for _, f := range files {
		fmt.Printf("read file: %s \n", f)
		content, err := ioutil.ReadFile(f)
		if err == nil {
			rtv[f] = string(content[:])
		} else {
			return rtv, err
		}
	}

	return rtv, nil
}
