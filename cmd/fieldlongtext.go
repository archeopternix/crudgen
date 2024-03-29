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

// fieldlongtextCmd represents the fieldlongtext command
var fieldlongtextCmd = &cobra.Command{
	Use:   "longtext",
	Short: "longtext field added to an entity",
	Long: `Adds a longtext field to an entity where you can set if the 
field is --required and define the maximum length.
`,
	Run: func(cmd *cobra.Command, args []string) {
		f := ast.Field{
			Name:     viper.GetString("name"),
			Kind:     "longtext",
			Required: viper.GetBool("required"),
			Length:   viper.GetInt("length"),
			Size:     viper.GetInt("columns"),
			Rows:     viper.GetInt("rows"),
		}

		if err := addField(viper.GetString("entity"), f); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	addCmd.AddCommand(fieldlongtextCmd)
	fieldlongtextCmd.Flags().StringP("name", "n", "", "name of the field")
	fieldlongtextCmd.Flags().StringP("entity", "e", "", "entity where the field will be added")
	fieldlongtextCmd.Flags().IntP("length", "l", 120, "maximum text length")
	fieldlongtextCmd.Flags().Int("columns", 80, "columns for textfield")
	fieldlongtextCmd.Flags().Int("rows", 4, "rows for textfield")
	fieldlongtextCmd.MarkFlagRequired("name")
	fieldlongtextCmd.MarkFlagRequired("entity")
	fieldlongtextCmd.Flags().Bool("required", false, "content for field is required to be accepted (to activate: --required)")

	viper.BindPFlag("name", fieldlongtextCmd.Flags().Lookup("name"))
	viper.BindPFlag("entity", fieldlongtextCmd.Flags().Lookup("entity"))
	viper.BindPFlag("length", fieldlongtextCmd.Flags().Lookup("length"))
	viper.BindPFlag("columns", fieldlongtextCmd.Flags().Lookup("columns"))
	viper.BindPFlag("rows", fieldlongtextCmd.Flags().Lookup("rows"))
	viper.BindPFlag("required", fieldlongtextCmd.Flags().Lookup("required"))
}
