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
	"crudgen/internal"
	"os"

	"github.com/go-git/go-git/v5"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates all or part of the application",
	Long:  `Generates all or part of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func cloneModulesRepository() error {
	if internal.CheckMkdir(viper.GetString("module-path")) == nil {
		// only when directory is empty
		_, err := git.PlainClone(viper.GetString("module-path"), false, &git.CloneOptions{
			URL:      viper.GetString("module-pkg"),
			Progress: os.Stdout,
		})
		return err
	}
	return nil
}
