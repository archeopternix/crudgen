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

// relationCmd represents the relation command
var relationCmd = &cobra.Command{
	Use:   "relation",
	Short: "adds an entity relation to the application",
	Long: `a relation will be added to the configuration. You can choose as 
relation type onetomany. As a flag source and target have to be submitted as 
the both entitites that are in a relation to each other`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add relation called")
	},
}

func init() {
	addCmd.AddCommand(relationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// relationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// relationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	relationCmd.Flags().String("source", "", "Name of the source (e.g. 1..) entity ")
	relationCmd.Flags().String("target", "", "Name of the target (e.g. ..n) entity ")
	relationCmd.Flags().String("type", "onetomany", "Type of relation (1..n = onetomany)")
}
