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
	"os"

	"github.com/spf13/cobra"
)

// addfieldCmd represents the addfield command
var addfieldCmd = &cobra.Command{
	Use:   "addfield",
	Short: "Adds fields to existing entitites",
	Long: `Adds fields to existing entitites. The mandatory sub-command reflects 
the type of field that has to be added.

Possible sub-commands:
Text, Password, Integer, Number, Boolean, Email, Tel, Longtext, Time, Lookup.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(addfieldCmd)

}
