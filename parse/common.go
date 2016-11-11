package parse

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/go-openapi/spec"
)

type Parser interface {
	parseComment(comment string) spec.PathItem
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

func readFiles(paths []string, ext string) []string {
	var rtv []string
	if ext == "" {
		return rtv
	}

	var files []string
	if paths != nil {
		for _, p := range paths {
			if info, err := os.Stat(p); os.IsNotExist(err) {
				log.Fatalf("path: %s not exist \n", p)
			} else if info.IsDir() {
				files = apend(files, recursiveFindFiles(p, ext))
			} else {
				if ext != path.Ext(p) {
					log.Printf("invalid ext: %s", path.Ext(p))
					continue
				}

			}

		}
	}
}
