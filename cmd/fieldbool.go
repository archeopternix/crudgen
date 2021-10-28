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

	"github.com/spf13/cobra"
)

// fieldboolCmd represents the fieldtext command
var fieldboolCmd = &cobra.Command{
	Use:   "boolean",
	Short: "adds a boolean field to an entity",
	Long: `Adds a boolean (true/false) field to an entity.
`,
	Run: func(cmd *cobra.Command, args []string) {
		addFieldBool()
	},
}

func init() {
	addCmd.AddCommand(fieldboolCmd)
	fieldboolCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the field")
	fieldboolCmd.Flags().StringVarP(&entity, "entity", "e", "", "Entity where the field will be added")
	fieldboolCmd.MarkFlagRequired("name")
	fieldboolCmd.MarkFlagRequired("entity")
}

func addFieldBool() {
	var a ast.Application

	if err := a.LoadFromYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := ast.Field{Name: name, Kind: "Boolean"}

	if err := a.AddFieldToEntity(entity, f); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.SaveToYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("New boolean field '%v' added to entity '%v'\n", name, entity)
}
