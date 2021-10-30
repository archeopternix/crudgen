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
var entity, kind, name, pkgname, source, target, typ, cfgpath, modulepath string
var required, label bool
var length, size, rows, step, min, max int

const (
	configfile     = ".crudgen"
	definitionfile = ".model"
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&name, "name", "n", "", "name of the application")
	initCmd.Flags().StringVar(&pkgname, "module-pkg", "https://github.com/archeopternix/crudgen-modules.git", "crudgen modules package in github.com")
	initCmd.Flags().StringVar(&modulepath, "module-path", "./modules/", "path where the modules are stored")
	initCmd.Flags().StringVar(&cfgpath, "cfgpath", "./config/", "path to config files")

	initCmd.MarkFlagRequired("name")
}

func createConfiguration() {
	// check if file already exists - call init only once

	if internal.FileExist(cfgpath+configfile) != nil {
		fmt.Println("Error: Init can be executed only once")
		os.Exit(1)
	}

	// create a directory when file is not found

	if err := internal.CheckMkdir(cfgpath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create Viper config file
	viper.Set("cfgpath", cfgpath)
	viper.Set("name", name)
	viper.Set("module-pkg", pkgname)
	viper.Set("module-path", modulepath)
	viper.SetConfigType("yaml")

	if err := viper.SafeWriteConfigAs(cfgpath + configfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Info: New config file created: ", cfgpath+configfile)

	a := ast.NewApplication(name)
	a.Name = name
	a.Config.PackageName = pkgname
	a.Config.DateFormat = "02.01.2006"
	a.Config.TimeFormat = "15:04:05.000"

	if err := a.SaveToYAML(cfgpath + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Info: New definition file created: ", cfgpath+definitionfile)

}
