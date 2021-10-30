// Package ast consists of the full AST (abstract syntax tree) which reflects
// the object structure consisting of Entities, Fields, Relations..

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

// relationCmd represents the relation command
var relationCmd = &cobra.Command{
	Use:   "relation",
	Short: "adds an entity relation to the application",
	Long: `The relation will be added to the configuration. You can choose as 
relation type onetomany. As a flag source and child entity names have to be submitted as 
the both entitites that are in a relation to each other`,
	Run: func(cmd *cobra.Command, args []string) {
		addRelation()
		fmt.Println("add relation called")
	},
}

func init() {
	addCmd.AddCommand(relationCmd)

	relationCmd.Flags().StringVarP(&source, "source", "s", "", "Name of the source (e.g. 1..) entity ")
	relationCmd.Flags().StringVarP(&target, "child", "c", "", "Name of the child (e.g. ..n) entity ")
	relationCmd.Flags().StringVar(&kind, "type", "onetomany", "Type of relation (e.g 1..n = onetomany)")
	relationCmd.MarkFlagRequired("child")
	relationCmd.MarkFlagRequired("target")

}

func addRelation() {
	var a ast.Application

	if err := a.LoadFromYAML(viper.GetString("cfgpath") + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.AddRelation(ast.Relation{Source: source, Child: target, Kind: kind}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := a.SaveToYAML(viper.GetString("cfgpath") + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("New relation '", source, target, typ, "' added to config file ")
}
