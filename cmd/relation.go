/*Package cmd is the command line interface to the CRUD Package generator


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
relation type onetomany. As a flag parent and child entity names have to be submitted as 
the both entitites that are in a relation to each other`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := addRelation(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("New relation '", source, target, typ, "' added to config file ")
	},
}

func init() {
	addCmd.AddCommand(relationCmd)

	relationCmd.Flags().StringVarP(&source, "parent", "s", "", "Name of the parent (e.g. 1..) entity ")
	relationCmd.Flags().StringVarP(&target, "child", "c", "", "Name of the child (e.g. ..n) entity ")
	relationCmd.Flags().StringVar(&kind, "type", "onetomany", "Type of relation (e.g 1..n = onetomany)")
	relationCmd.MarkFlagRequired("parent")
	relationCmd.MarkFlagRequired("child")

}

func addRelation() error {
	a, err := ast.NewFromYAMLFile(viper.GetString("cfgpath") + definitionfile)
	if err != nil {
		return err
	}

	if err := a.AddRelation(ast.Relation{Parent: source, Child: target, Kind: kind}); err != nil {
		return err

	}

	if err := a.SaveToYAML(viper.GetString("cfgpath") + definitionfile); err != nil {
		return err
	}
	return nil
}
