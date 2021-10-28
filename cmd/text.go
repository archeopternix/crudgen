/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Creates a text field",
	Long: `A text field is created. Optional flags are
`,
	Run: func(cmd *cobra.Command, args []string) {
		addTextField()
	},
}

func init() {
	addfieldCmd.AddCommand(textCmd)
	textCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the field")
	textCmd.Flags().StringVarP(&entity, "entity", "e", "", "Entity where the field will be added")
	textCmd.MarkFlagRequired("name")
	textCmd.MarkFlagRequired("entity")
	textCmd.Flags().BoolVarP(&required, "required", "", false, "Content for field is required to be accepted (to activate: --required)")
	textCmd.Flags().BoolVarP(&label, "label", "", false, "This field will be used as a label for drop down fields (to activate: --label)")

}

func addTextField() {
	var a ast.Application

	if err := a.LoadFromYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := ast.Field{Name: name, Kind: "Text", Required: required, IsLabel: label}

	if err := a.AddFieldToEntity(entity, f); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.SaveToYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("New field '", name, "' added to entity '", entity, "'")
}
