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
	"io/fs"
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

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.PersistentFlags().String("pkg-name", "", "provide the package name")
	viper.BindPFlag("pkg-name", initCmd.PersistentFlags().Lookup("pkg-name"))

}

func createConfiguration() {
	// check if file already exists - call init only once
	file, err := os.Open("./config/.crudgen")
	defer file.Close()
	if err == nil {
		fmt.Println("Init can be executed only once")
		os.Exit(1)
	}

	err = os.Mkdir("./config/", fs.FileMode(0755))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigType("yaml")
	err = viper.SafeWriteConfigAs("./config/.crudgen")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Init, new config file created: ", viper.ConfigFileUsed())

}
