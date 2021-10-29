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
	"math"
	"os"

	"github.com/spf13/cobra"
)

// fieldtextCmd represents the fieldtext command
var fieldintegerCmd = &cobra.Command{
	Use:   "integer",
	Short: "adds an integer field to an entity",
	Long: `Adds an integer field to an entity where you can set the 'min', 'max' value 
	that is allowed to enter. The standard 'step' between values is 1 (means integer) but this can 
	be changed by setting the 'step' flag
`,
	Run: func(cmd *cobra.Command, args []string) {
		addFieldInteger()
	},
}

func init() {
	addCmd.AddCommand(fieldintegerCmd)
	fieldintegerCmd.Flags().StringVarP(&name, "name", "n", "", "name of the field")
	fieldintegerCmd.Flags().StringVarP(&entity, "entity", "e", "", "entity where the field will be added")
	fieldintegerCmd.Flags().IntVarP(&step, "step", "", 1, "step between values")
	fieldintegerCmd.Flags().IntVarP(&min, "min", "", math.MinInt, "minimum value for field")
	fieldintegerCmd.Flags().IntVarP(&max, "max", "", math.MaxInt, "maximum value for field")
	fieldintegerCmd.Flags().IntVarP(&length, "length", "l", 12, "maximum text length")
	fieldintegerCmd.Flags().IntVarP(&size, "size", "s", 80, "size of the entry field")

	fieldintegerCmd.MarkFlagRequired("name")
	fieldintegerCmd.MarkFlagRequired("entity")
}

func addFieldInteger() {
	var a ast.Application

	if err := a.LoadFromYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := ast.Field{Name: name, Kind: "Integer", Step: step, Min: min, Max: max, Length: length, Size: size}

	if err := a.AddFieldToEntity(entity, f); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.SaveToYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("New integer field '%v' added to entity '%v'\n", name, entity)
}
