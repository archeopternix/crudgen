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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generatecoreCmd represents the generatecore command
var generatecoreCmd = &cobra.Command{
	Use:   "core",
	Short: "generates the core components",
	Long:  `Generates the core components.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := cloneModulesRepository(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generateCore(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.AddCommand(generatecoreCmd)

}

func generateCore() error {
	c := NewGenerator()
	if err := c.AddModule(viper.GetString(""module-path"")+ "application/app.yaml"); err != nil {
			log.Fatalf("ERROR: %v", err)
		}
}
