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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generaterepositoryCmd represents the generaterepository command
var generaterepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "repository for data storage",
	Long:  `Creates a respository and associated tests.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := cloneModulesRepository(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var modules []string
		if repo == "Mock" {
			modules = []string{
				"databasetest/databasetest.yaml",
				"mockdatabase/mockdatabase.yaml",
			}
		} else {
			modules = []string{
				"databasetest/databasetest.yaml",
				"sqldatabase/database.yaml.yaml",
			}
		}

		viper.Set("database", Environment{
			Instance: "DEV",
			Database: "postgres",
			Host:     "localhost",
			Port:     5432,
			User:     "admin",
			Password: "****",
			Dbname:   "my-db"})

		if err := runModuleCreation(modules); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.AddCommand(generaterepositoryCmd)
	generaterepositoryCmd.Flags().StringVarP(&repo, "repository", "r", "Mock", "selection which repository to choose [Mock, SQL]")

}
