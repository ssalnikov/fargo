// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"path/filepath"
	"strings"

	"github.com/gigovich/fargo/core/parser"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "run codegeneration for all application modules",
	Long:  `First of all this command regenerate model definions`,
	Run: func(cmd *cobra.Command, args []string) {
		// walt to current working dir, and parse all go files to run codegeneration
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}

		filepath.Walk(dir, func(filePath string, info os.FileInfo, err error) error {
			split := strings.Split(info.Name(), ".")
			if info.IsDir() || split[len(split)-1] != "go" {
				return nil
			}

			models, err := parser.New(filePath).Parse()
			if err != nil {
				return err
			}

			fmt.Printf("%+v", models)

			return nil
		})
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
