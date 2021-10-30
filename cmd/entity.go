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
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// entityCmd represents the entity command
var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "adds an entity to the application",
	Long: `An entity will be added to the configuration. The default type is a
normal 'entity' that holds fields, it is necessary to create fields and add 
them to the entity configuration.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := addEntity(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("New entity '", name, "' added to config file ")
	},
}

func init() {
	addCmd.AddCommand(entityCmd)
	entityCmd.Flags().StringVarP(&name, "name", "n", "", "name of the entity")
	entityCmd.Flags().StringVarP(&kind, "type", "t", "default", "type of the entity to be created")
	entityCmd.MarkFlagRequired("name")
}

func addEntity() error {
	a, err := ast.NewFromYAMLFile(viper.GetString("cfgpath") + definitionfile)
	if err != nil {
		return err
	}

	if err := a.AddEntity(ast.Entity{Name: name, Kind: kind}); err != nil {
		return err
	}

	if err := a.SaveToYAML(viper.GetString("cfgpath") + definitionfile); err != nil {
		return err
	}
	return nil
}
