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

// fieldlookupCmd represents the fieldlookup command
var fieldlookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "adds a lookup field to an entity",
	Long: `Adds a lookup field to an entity. The name of the lookupfield needs to 
be the name of the corresponding entity
`,
	Run: func(cmd *cobra.Command, args []string) {
		addFieldLookup()
	},
}

func init() {
	addCmd.AddCommand(fieldlookupCmd)
	fieldlookupCmd.Flags().StringVarP(&name, "name", "n", "", "name of the field equals the coresponding entity")
	fieldlookupCmd.Flags().StringVarP(&entity, "entity", "e", "", "entity where the field will be added")

	fieldlookupCmd.MarkFlagRequired("name")
	fieldlookupCmd.MarkFlagRequired("entity")
}

func addFieldLookup() {
	var a ast.Application

	if err := a.LoadFromYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := ast.Field{Name: name, Kind: "Lookup"}

	if err := a.AddFieldToEntity(entity, f); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.SaveToYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("New lookup field '%v' added to entity '%v'\n", name, entity)
}
