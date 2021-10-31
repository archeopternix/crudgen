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
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start and run the application",
	Long: `The application will be generated (generate all) and
the main() will be executed. If there is a webserver needed the option flag port 
could be used - as a default port 8080 is used`,
	Run: func(cmd *cobra.Command, args []string) {

		// initialize or refresh go.mod
		prg := "go"
		arg1 := "build"

		command := exec.Command(prg, arg1)
		if err := command.Run(); err != nil {
			fmt.Printf("Error in build: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("weberver started. Hit CTRL+C to stop")
		command = exec.Command("_test.exe")
		err := command.Run()
		if err != nil {
			fmt.Printf("Error in run: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().Int16("port", 8080, "Port to be used by webserver")
}
