// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/koolay/scmt/parse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var langMap = map[string]string{"php": ".php"}

func NewParser(lang string) (parse.Parser, error) {
	switch lang {
	case "php":
		return new(parse.PhpParser), nil
	default:
		return nil, errors.New("do not support language:" + lang)
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Output json of swagger from special file or directinary",
	Long:  `Output json of swagger from special file or directinary.`,
	Run: func(cmd *cobra.Command, args []string) {
		sources := viper.Get("sources").([]string)
		if len(sources) == 0 {
			fmt.Println("No sources.")
			return
		}
		lang := viper.Get("lang").(string)
		ext := langMap[lang]
		if ext == "" {
			log.Fatal("do not support the language: " + lang)
			return
		} else {
			parser, err := NewParser(lang)
			if err == nil {
				log.Println("start read " + ext + "...")
				for _, source := range sources {
					log.Println("read " + source)
					parser.Parse(source)
				}
			} else {
				log.Fatal(err.Error())
			}
		}

		fmt.Println("create called")
	},
}

func init() {
	RootCmd.AddCommand(createCmd)

}
