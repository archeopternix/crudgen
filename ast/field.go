// Package ast consists of the full AST (abstract syntax tree) which reflects
// the object structure consisting of Entities, Fields, Relations..

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
package ast

// Field is the definition for every single attribute within an entity.
// Object is empty except in case type=lookup or child keeps the name of the Object
type Field struct {
	Name     string `yaml:"name"`
	Kind     string `yaml:"kind"` // string, int, bool, lookup, tel, email
	Required bool   `yaml:"required,omitempty"`
	IsLabel  bool   `yaml:"islabel,omitempty"` // when true is the shown text for select boxes
	Object   string `yaml:"object,omitempty"`  // for lookup, child relations - mappingtable for many2many relations
	Length   int    `yaml:"length,omitempty"`
	Size     int    `yaml:"size,omitempty"` // for textarea size = cols

	Step int `yaml:"step,omitempty"` //for Number fields
	Min  int `yaml:"min,omitempty"`  //for Number fields
	Max  int `yaml:"max,omitempty"`  //for Number fields

	Rows int `yaml:"rows,omitempty"` //for textarea
}
