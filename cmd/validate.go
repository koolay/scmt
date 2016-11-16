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
	"fmt"
	"log"

	swaggererrors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"github.com/spf13/cobra"
)

var (
	url string
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate swagger content",
	Long:  `validate swagger content from json,yml`,
	Run: func(cmd *cobra.Command, args []string) {

		if url == "" {
			fmt.Println("The validate command requires the swagger document url to be specified")
			return
		}

		specDoc, err := loads.Spec(url)
		if err != nil {
			log.Fatalln(err)
		}

		result := validate.Spec(specDoc, strfmt.Default)
		if result == nil {
			fmt.Printf("The swagger spec at %q is valid against swagger specification %s\n", url, specDoc.Version())
			return
		} else {
			str := fmt.Sprintf("The swagger spec at %q is invalid against swagger specification %s. see errors :\n", url, specDoc.Version())
			for _, desc := range result.(*swaggererrors.CompositeError).Errors {
				str += fmt.Sprintf("- %s\n", desc)
			}
			fmt.Printf(str)
		}

	},
}

func init() {
	RootCmd.AddCommand(validateCmd)

	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	validateCmd.Flags().StringVar(&url, "url", "", "swagger url")

}
