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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fieldpasswordCmd represents the fieldtext command
var fieldpasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "password field added to an entity",
	Long: `Adds a password field to an entity where you can set if the field is --required 
and define the maximum length. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		f := ast.Field{
			Name:     viper.GetString("name"),
			Kind:     "password",
			Required: viper.GetBool("required"),
			Length:   viper.GetInt("length"),
			Size:     viper.GetInt("size"),
		}

		if err := addField(viper.GetString("entity"), f); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	addCmd.AddCommand(fieldpasswordCmd)
	fieldpasswordCmd.Flags().StringP("name", "n", "", "name of the field")
	fieldpasswordCmd.Flags().StringP("entity", "e", "", "entity where the field will be added")
	fieldpasswordCmd.Flags().IntP("length", "l", 20, "maximum text length")
	fieldpasswordCmd.Flags().IntP("size", "s", 80, "size of the entry field")
	fieldpasswordCmd.MarkFlagRequired("name")
	fieldpasswordCmd.MarkFlagRequired("entity")
	fieldpasswordCmd.Flags().Bool("required", false, "content for field is required to be accepted (to activate: --required)")

	viper.BindPFlag("name", fieldpasswordCmd.Flags().Lookup("name"))
	viper.BindPFlag("entity", fieldpasswordCmd.Flags().Lookup("entity"))
	viper.BindPFlag("length", fieldpasswordCmd.Flags().Lookup("length"))
	viper.BindPFlag("size", fieldpasswordCmd.Flags().Lookup("size"))
	viper.BindPFlag("required", fieldpasswordCmd.Flags().Lookup("required"))
}
