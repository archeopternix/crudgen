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
)

// fieldphoneCmd represents the fieldtext command
var fieldphoneCmd = &cobra.Command{
	Use:   "phone",
	Short: "phone field added to an entity",
	Long: `Adds a phone field to an entity where you can set if the field is --required 
and define the maximum length. Length=-1 means no restriction
`,
	Run: func(cmd *cobra.Command, args []string) {
		f := ast.Field{Name: name, Kind: "phone", Required: required, Length: length, Size: size}

		if err := addField(f); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	addCmd.AddCommand(fieldphoneCmd)
	fieldphoneCmd.Flags().StringVarP(&name, "name", "n", "", "name of the field")
	fieldphoneCmd.Flags().StringVarP(&entity, "entity", "e", "", "entity where the field will be added")
	fieldphoneCmd.Flags().IntVarP(&length, "length", "l", 20, "maximum text length")
	fieldphoneCmd.Flags().IntVarP(&size, "size", "s", 80, "size of the entry field")
	fieldphoneCmd.MarkFlagRequired("name")
	fieldphoneCmd.MarkFlagRequired("entity")
	fieldphoneCmd.Flags().BoolVarP(&required, "required", "", false, "content for field is required to be accepted (to activate: --required)")
}
