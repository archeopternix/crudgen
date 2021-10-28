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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialisation of the AST configuration in the target directory",
	Long: `Basic setup of the AST configuration in the target directory. 
Configuration files will be created with default data set.`,
	Run: func(cmd *cobra.Command, args []string) {
		createConfiguration()
	},
}

// Package local variables used to capture commands
var kind, name, pkgname, source, target, typ string

const (
	configpath     = "./config/"
	configfile     = ".crudgen"
	definitionfile = ".model"
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the application")
	initCmd.Flags().StringVar(&pkgname, "pkg-name", "", "Package name of the root package (e.g. github.com/abc)")
	initCmd.MarkFlagRequired("name")
	initCmd.MarkFlagRequired("pkg-name")
}

func createConfiguration() {
	// check if file already exists - call init only once

	if internal.FileExist(configpath+configfile) != nil {
		fmt.Println("Error: Init can be executed only once")
		os.Exit(1)
	}

	// create a directory when file is not found

	if err := internal.CheckMkdir(configpath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create Viper config file
	viper.Set("name", name)
	viper.Set("pkg-name", pkgname)
	viper.SetConfigType("yaml")

	if err := viper.SafeWriteConfigAs(configpath + configfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Info: New config file created: ", configpath+configfile)

	a := ast.NewApplication(name)
	a.Config.PackageName = pkgname

	if err := a.SaveToYAML(configpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Info: New definition file created: ", configpath+definitionfile)

}
