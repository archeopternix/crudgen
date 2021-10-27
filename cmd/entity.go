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
	"fmt"

	"github.com/spf13/cobra"
)

// entityCmd represents the entity command
var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "adds an entity to the application",
	Long: `an entity will be added to the configuration. The default type is a
normal 'entity' that holds fields, it is necessary to create fields and add 
them to the entity configuration.

A special entity type is 'lookup' which could populate drop down fields.`,
	Run: func(cmd *cobra.Command, args []string) {
		addEntity()
		fmt.Println("add entity called")
	},
}

var ename, kind string

func init() {
	addCmd.AddCommand(entityCmd)
	entityCmd.Flags().StringVarP(&ename, "name", "n", "", "Name of the entity")
	entityCmd.Flags().StringVarP(&kind, "type", "t", "default", "Type of the entity to be created (default or lookup")
	entityCmd.MarkFlagRequired("name")
}

func addEntity() {
	/*	if app == nil {
			fmt.Println("ERROR: initialize application first.'")
			os.Exit(1)
		}

		if _, ok := app.Entities[ename]; ok {
			fmt.Println("ERROR: Entity already exists: '", ename, "'")
			os.Exit(1)
		}
		app.Entities[ename] = ast.Entity{Name: ename, Kind: kind}

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	*/
	fmt.Println("New entity ", ename, " added to config file ")
}
