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

	//	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crudgen",
	Short: "A brief description of your application",
	Long: `Generator for a web based CRUD application/API with selectable frontends and backends. 
CRUDgen uses an AST tree that will build up based on configuration files in YAML. 
You will be provided with interactive shell commends that helps you building up the 
application stack.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/crudgen.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		fmt.Println("config file set to: '", cfgFile, "'")
	} else {

		// Search config in home directory with name ".crudgen" (without extension).
		viper.AddConfigPath("./config/")
		viper.SetConfigName("crudgen")
		viper.SetConfigType("yaml")
	}

	/*
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("read of config file: :", err)
			os.Exit(1)
			// fmt.Println("Info: Using config file:", viper.ConfigFileUsed())
		}
	*/
	viper.AutomaticEnv() // read in environment variables that match

}
