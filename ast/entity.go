/*Package ast consists of the full AST (abstract syntax tree) which reflects
the object structure consisting of Entities, Fields, Relations..

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
package ast

import (
	"crudgen/internal"
	"fmt"
)

// Entity relates to a database table and holds the field definitions
type Entity struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
	Kind   string  `yaml:"type,omitempty"` // 0..default, 1..lookup 2..many2many
}

// EntityCheckForErrors checks an entity for errors.
// 'Name' has to be longer than 3 characters without whitespaces
// 'Kind' is default or lookup
func EntityCheckForErrors(e Entity) error {
	if len(e.Name) < 4 {
		return fmt.Errorf("Entity needs a unique name (min 3 characters): '%v'", e.Name)
	}

	if !internal.IsLetter(e.Name) {
		return fmt.Errorf("Entity must contain only letters [a-zA-Z0-9]: '%v'", e.Name)
	}

	switch e.Kind {
	case "default":
	case "lookup":

	default:
		return fmt.Errorf("Missing or unknown entity type: '%v'", e.Kind)
	}
	return nil
}
