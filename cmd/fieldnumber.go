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
var fieldnumberCmd = &cobra.Command{
	Use:   "number",
	Short: "number field added to an entity",
	Long: `Adds a number field to an entity. Numbers are any floating point values
`,
	Run: func(cmd *cobra.Command, args []string) {
		f := ast.Field{
			Name:     viper.GetString("name"),
			Kind:     "number",
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
	addCmd.AddCommand(fieldnumberCmd)
	fieldnumberCmd.Flags().StringP("name", "n", "", "name of the field")
	fieldnumberCmd.Flags().StringP("entity", "e", "", "entity where the field will be added")
	fieldnumberCmd.Flags().IntP("length", "l", 30, "maximum text length")
	fieldnumberCmd.Flags().IntP("size", "s", 80, "size of the entry field")
	fieldnumberCmd.MarkFlagRequired("name")
	fieldnumberCmd.MarkFlagRequired("entity")

	viper.BindPFlag("name", fieldnumberCmd.Flags().Lookup("name"))
	viper.BindPFlag("entity", fieldnumberCmd.Flags().Lookup("entity"))
	viper.BindPFlag("length", fieldnumberCmd.Flags().Lookup("length"))
	viper.BindPFlag("size", fieldnumberCmd.Flags().Lookup("size"))
}
