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
	"math"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fieldtextCmd represents the fieldtext command
var fieldintegerCmd = &cobra.Command{
	Use:   "integer",
	Short: "integer field added to an entity",
	Long: `Adds an integer field to an entity where you can set the 'min', 'max' value 
	that is allowed to enter. The standard 'step' between values is 1 (means integer) but this can 
	be changed by setting the 'step' flag
`,
	Run: func(cmd *cobra.Command, args []string) {
		f := ast.Field{
			Name:   viper.GetString("name"),
			Kind:   "integer",
			Length: viper.GetInt("length"),
			Size:   viper.GetInt("size"),
			Step:   viper.GetInt("step"),
			Min:    viper.GetInt("min"),
			Max:    viper.GetInt("max"),
		}

		if err := addField(viper.GetString("entity"), f); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	addCmd.AddCommand(fieldintegerCmd)
	fieldintegerCmd.Flags().StringP("name", "n", "", "name of the field")
	fieldintegerCmd.Flags().StringP("entity", "e", "", "entity where the field will be added")
	fieldintegerCmd.Flags().Int("step", 1, "step between values")
	fieldintegerCmd.Flags().Int("min", math.MinInt, "minimum value for field")
	fieldintegerCmd.Flags().Int("max", math.MaxInt, "maximum value for field")
	fieldintegerCmd.Flags().IntP("length", "l", 12, "maximum text length")
	fieldintegerCmd.Flags().IntP("size", "s", 80, "size of the entry field")
	fieldintegerCmd.MarkFlagRequired("name")
	fieldintegerCmd.MarkFlagRequired("entity")

	viper.BindPFlag("name", fieldintegerCmd.Flags().Lookup("name"))
	viper.BindPFlag("entity", fieldintegerCmd.Flags().Lookup("entity"))
	viper.BindPFlag("length", fieldintegerCmd.Flags().Lookup("length"))
	viper.BindPFlag("size", fieldintegerCmd.Flags().Lookup("size"))
	viper.BindPFlag("step", fieldintegerCmd.Flags().Lookup("step"))
	viper.BindPFlag("min", fieldintegerCmd.Flags().Lookup("min"))
	viper.BindPFlag("max", fieldintegerCmd.Flags().Lookup("max"))

}
