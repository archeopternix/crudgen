// Package cmd is the command line interface to the CRUD Package generator

/*
Copyright © 2021 Andreas<DOC>Eisner <andreas.eisner@kouri.cc>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"crudgen/ast"
	"crudgen/internal"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateviewCmd represents the generateview command
var generateviewCmd = &cobra.Command{
	Use:   "view",
	Short: "view components will be generated",
	Long:  `A webserver including html templates are generated`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := cloneModulesRepository(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generateView(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.AddCommand(generateviewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateView() error {
	gen := internal.NewGenerator()

	if err := gen.ModuleFromYAML(viper.GetString("module-path") + "view/view.yaml"); err != nil {
		return err
	}

	a, err := ast.NewFromYAMLFile(viper.GetString("cfgpath") + definitionfile)
	if err != nil {
		return err
	}

	gen.Worker = ast.NewGeneratorWorker(a)

	if err := gen.GenerateAll(); err != nil {
		return err
	}
	return nil
}