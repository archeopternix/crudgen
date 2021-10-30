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
	"crudgen/internal"
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
		if err := generateRepo(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.AddCommand(generaterepositoryCmd)
	generaterepositoryCmd.Flags().StringVarP(&repo, "repository", "r", "Mock", "selection which repository to choose [Mock, SQL]")

}

type Environment struct {
	Instance string `yaml:"instance"` // production, development, testing
	Database string `yaml:"database"` // postgres,mysql...
	Host     string `yaml:"host"`     // localhost or IP address
	Port     int    `yaml:"port"`     //5432
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func generateRepo() error {
	gen := internal.NewGenerator()

	if err := gen.ModuleFromYAML(viper.GetString("module-path") + "databasetest/databasetest.yaml"); err != nil {
		return err
	}
	if repo == "Mock" {
		if err := gen.ModuleFromYAML(viper.GetString("module-path") + "mockdatabase/mockdatabase.yaml"); err != nil {
			return err
		}
		fmt.Println("Mock repo installed")
	}
	if repo == "SQL" {
		if err := gen.ModuleFromYAML(viper.GetString("module-path") + "sqldatabase/database.yaml"); err != nil {
			return err
		}

		viper.Set("database", Environment{
			Instance: "DEV",
			Database: "postgres",
			Host:     "localhost",
			Port:     5432,
			User:     "admin",
			Password: "****",
			Dbname:   "my-db"})

		fmt.Println("SQL repo installed")
	}

	a, err := ast.NewFromYAMLFile(viper.GetString("cfgpath") + definitionfile)
	if err != nil {
		return err
	}

	gen.Worker = ast.NewGeneratorWorker(a)

	if err := gen.GenerateAll(); err != nil {
		return err
	}
	return nil
}
