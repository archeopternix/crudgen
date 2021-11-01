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

// fieldtextCmd represents the fieldtext command
var fieldtextCmd = &cobra.Command{
	Use:   "text",
	Short: "text field added to an entity",
	Long: `Adds a text field to an entity where you can set if the field is --required 
or used as a --label in drop down select boxes and define the maximum length. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		f := ast.Field{
			Name:     viper.GetString("name"),
			Kind:     "text",
			Required: viper.GetBool("required"),
			IsLabel:  viper.GetBool("label"),
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
	addCmd.AddCommand(fieldtextCmd)
	fieldtextCmd.Flags().StringP("name", "n", "", "name of the field")
	fieldtextCmd.Flags().StringP("entity", "e", "", "entity where the field will be added")
	fieldtextCmd.MarkFlagRequired("name")
	fieldtextCmd.MarkFlagRequired("entity")
	fieldtextCmd.Flags().IntP("length", "l", 120, "maximum text length")
	fieldtextCmd.Flags().IntP("size", "s", 80, "size of the entry field")
	fieldtextCmd.Flags().Bool("required", false, "content for field is required to be accepted (to activate: --required)")
	fieldtextCmd.Flags().Bool("label", false, "field will be used as a label for drop down fields (to activate: --label)")

	viper.BindPFlag("name", fieldtextCmd.Flags().Lookup("name"))
	viper.BindPFlag("entity", fieldtextCmd.Flags().Lookup("entity"))
	viper.BindPFlag("length", fieldtextCmd.Flags().Lookup("length"))
	viper.BindPFlag("size", fieldtextCmd.Flags().Lookup("size"))
	viper.BindPFlag("required", fieldtextCmd.Flags().Lookup("required"))
	viper.BindPFlag("label", fieldtextCmd.Flags().Lookup("label"))
}
