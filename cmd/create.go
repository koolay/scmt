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
	"io/ioutil"
	"log"
	"os"

	"github.com/koolay/scmt/parse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (

	// output to
	OUTPUT_STDOUT = ""
	OUTPUT_JSON   = "json"
	OUTPUT_API    = "api"
	OUTPUT_YML    = "yml"
)

var (
	sources []string
	cfgFile string
	// language php, python ..
	lang string
	// where to output
	outputs []string
	// name of swagger
	name    string
	version string
	// if output is api
	httpHeaders []string
)

var langMap = map[string]string{"php": ".php", "python": ".py"}

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
		}
		swaggerParser := parse.NewParser(lang)
		for _, source := range sources {
			log.Println("read from: " + source)
			files, err := parse.FindFiles(source, ext)
			if err != nil {
				fmt.Printf("read files failed. source: %s", source)
				panic(err)
			}

			for _, f := range files {

				fmt.Printf("read file: %s \n", f)
				content, err := ioutil.ReadFile(f)
				if err != nil {
					panic(err)
				}
				sourceCode := string(content[:])
				swaggerParser.Parse(sourceCode)
			}
		}

		outputer := OutPuter{}
		outputer.OutputFlags = outputs
		outputer.Swagger = swaggerParser.Swagger
		err := outputer.Output()
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	RootCmd.AddCommand(createCmd)

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	createCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scmt.yaml)")
	createCmd.PersistentFlags().StringArrayVarP(&sources, "sources", "s", nil, "full path of special directory or file")
	createCmd.PersistentFlags().StringVarP(&lang, "language", "l", "", "language, php,pytho,go etc.")
	createCmd.PersistentFlags().StringVar(&name, "name", "", "name of swagger project.")
	createCmd.PersistentFlags().StringVar(&version, "version", "", "version of swagger project.")
	createCmd.PersistentFlags().StringArrayVarP(&outputs, "output", "o", []string{"json"}, `Where to output, can be json/api/yml.
	eg:
	output to a json file: -o /home/koolay/swagger.json
	output to a yml file: -o /home/koolay/swagger.yml
	output to POST an api: -o http://myhost.com/swagger
	`)
	createCmd.PersistentFlags().StringArrayVarP(&httpHeaders, "headers", "H", nil, `http headers.
	eg:
	-H Authorization="Bearer mytoken"
	`)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".scmt") // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	if err := verifyArgs(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	viper.Set("sources", sources)
	viper.Set("lang", lang)
	viper.Set("output", outputs)
	viper.Set("name", name)
	viper.Set("version", version)
	viper.Set("headers", httpHeaders)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func verifyArgs() error {

	if sources == nil {
		return errors.New("Miss args of sources")
	}

	if lang == "" {
		return errors.New("Miss args of lang")
	}
	return nil
}
