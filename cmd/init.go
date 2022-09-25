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
	"crudgen/internal"
	"fmt"

	"os"
	"path/filepath"

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

/* Package local variables used to capture commands
var entity, kind, name, pkgname, source, target, typ, cfgpath, modulepath, repo string
var required, label bool
var length, size, rows, step, min, max int
*/

const (
	definitionfile = "model.yaml"
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("name", "n", "", "name of the application")
	initCmd.Flags().String("module-pkg", "https://github.com/archeopternix/crudgen-modules.git", "crudgen modules package in github.com")
	initCmd.Flags().String("module-path", "./modules/", "path where the modules are stored")
	initCmd.Flags().String("cfgpath", "./config/", "path to config files")
	initCmd.Flags().String("cfgfile", "crudgen.yaml", "filename for config files")
	initCmd.MarkFlagRequired("name")

	viper.BindPFlag("name", initCmd.Flags().Lookup("name"))
	viper.BindPFlag("module-pkg", initCmd.Flags().Lookup("module-pkg"))
	viper.BindPFlag("module-path", initCmd.Flags().Lookup("module-path"))
	viper.BindPFlag("cfgpath", initCmd.Flags().Lookup("cfgpath"))
	viper.BindPFlag("cfgfile", initCmd.Flags().Lookup("cfgfile"))

}

func createConfiguration() {
	// check if file already exists - call init only once

	if internal.FileExist(viper.GetString("cfgpath")+viper.GetString("cfgfile")) != nil {
		fmt.Println("Error: Init can be executed only once")
		os.Exit(1)
	}

	// create a directory when file is not found
	fmt.Println("PATH: ", viper.GetString("cfgpath"))
	if err := internal.CheckMkdir(viper.GetString("cfgpath")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create Viper config file
	//	viper.Set("cfgpath", cfgpath)
	//viper.Set("name", name)
	//viper.Set("module-pkg", pkgname)
	//viper.Set("module-path", modulepath)
	viper.SetConfigType("yaml")

	if err := viper.SafeWriteConfigAs(viper.GetString("cfgpath") + viper.GetString("cfgfile")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Info: New config file created: ", viper.GetString("cfgpath")+viper.GetString("cfgfile"))

	a := ast.NewApplication(viper.GetString("name"))
	a.Name = viper.GetString("name")
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	a.Config.PackageName = filepath.Base(path)
	a.Config.DateFormat = "02.01.2006"
	a.Config.TimeFormat = "15:04:05.000"
	a.Config.CurrencySymbol = "€"
	a.Config.DecimalSeparator = ","
	a.Config.ThousandSeparator = "."

	if err := a.SaveToYAML(viper.GetString("cfgpath") + definitionfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Info: New definition file created: ", viper.GetString("cfgpath")+definitionfile)

}
