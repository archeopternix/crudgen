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
	"os/exec"

	"github.com/spf13/cobra"
)

// generateallCmd represents the generateall command
var generateallCmd = &cobra.Command{
	Use:   "all",
	Short: "generates the full application",
	Long:  `Generates the full application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := cloneModulesRepository(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		modules := []string{
			"application/app.yaml",
			"model/models.yaml",
			"databasetest/databasetest.yaml",
			"mockdatabase/mockdatabase.yaml",
			"view/view.yaml",
		}
		if err := runModuleCreation(modules); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		// initialize or refresh go.mod
		prg := "go"
		arg1 := "build"

		command := exec.Command(prg, arg1)
		if err := command.Run(); err != nil {
			fmt.Printf("Error in build: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.AddCommand(generateallCmd)

}
