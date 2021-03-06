/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	model "github.com/archeopternix/crudgen/crudgen/model"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new application with a default landingpage",
	Long: `generates the basic logic to show a generic landing page, a config 
	file containing basic information, echo router with a default route '/', `,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println(arg)
		}
		// create sample data
		e, err := model.App.NewEntity("Sample1", model.Regular)
		if err != nil {
			log.Println(err)
		}

		f1 := model.Field{Name: "ID", FieldType: model.Integer, Required: true}
		e.AddField(&f1)
		f2 := model.Field{Name: "Name", FieldType: model.String}
		e.AddField(&f2)

		_, err = model.App.NewEntity("Sample2", model.Regular)
		if err != nil {
			log.Fatal(err)
		}

		_, err = model.App.NewRelation("Sample1", "Sample2", model.One2many)
		if err != nil {
			log.Fatal(err)
		}

		ds, er := model.NewYAMLDatastore("abc.yaml")
		if er != nil {
			log.Fatal(er)
		}
		ds.SaveAllData(model.App)
		fmt.Println("create called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().StringVarP(&Application, "application", "a", "", "application name - needed when no config file in place")
	createCmd.Flags().StringVarP(&Path, "path", "p", "", "target path of the generated application - needed when no config file in place")

	rootCmd.AddCommand(createCmd)
}
