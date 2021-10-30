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
	"fmt"

	"github.com/spf13/cobra"
)

// generatemodelCmd represents the generatemodel command
var generatemodelCmd = &cobra.Command{
	Use:   "model",
	Short: "generates the model structs",
	Long:  `Generates the model structs.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := cloneModulesRepository(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generateModel(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.AddCommand(generatemodelCmd)

}

func generateModel() error {
	gen := internal.NewGenerator()

	if err := gen.ModuleFromYAML(viper.GetString("module-path") + "model/models.yaml"); err != nil {
		return err
	}
	if err := gen.ModuleFromYAML(viper.GetString("module-path") + "mockdatabase/mockdatabase.yaml"); err != nil {
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
