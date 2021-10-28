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

// fieldemailCmd represents the fieldtext command
var fieldemailCmd = &cobra.Command{
	Use:   "email",
	Short: "adds a e-mail field to an entity",
	Long: `Adds a e-mail field to an entity where you can set if the field is --required 
and define the maximum length. Length=-1 means no restriction
`,
	Run: func(cmd *cobra.Command, args []string) {
		addFieldEmail()
	},
}

func init() {
	addCmd.AddCommand(fieldemailCmd)
	fieldemailCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the field")
	fieldemailCmd.Flags().StringVarP(&entity, "entity", "e", "", "Entity where the field will be added")
	fieldemailCmd.Flags().IntVarP(&length, "length", "l", -1, "Maximum text length (-1 .. means no restriction)")
	fieldemailCmd.MarkFlagRequired("name")
	fieldemailCmd.MarkFlagRequired("entity")
	fieldemailCmd.Flags().BoolVarP(&required, "required", "", false, "Content for field is required to be accepted (to activate: --required)")
}

func addFieldEmail() {
	var a ast.Application

	if err := a.LoadFromYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := ast.Field{Name: name, Kind: "Email", Required: required, Length: length}

	if err := a.AddFieldToEntity(entity, f); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.SaveToYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("New e-mail field '%v' added to entity '%v'\n", name, entity)
}
