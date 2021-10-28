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

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// Application holds all information and configuration for the AST and consists
// of Entitites, Relations and the Configuration necessary for template generation
type Application struct {
	Name      string              `yaml:"name"`
	Entities  map[string]Entity   `yaml:"entities"`  //Entity
	Relations map[string]Relation `yaml:"relations"` //Relation
	Config    struct {
		PackageName string `yaml:"packagename"`
	}
}

// NewApplication creates an new Application instance
func NewApplication(name string) *Application {
	app := new(Application)
	app.Name = name
	app.Entities = make(map[string]Entity)
	app.Relations = make(map[string]Relation)
	return app
}

// AddEntity adds an new entity to the AST and checks if Entity with this name already exists
// or name is too short
func (a *Application) AddEntity(e Entity) error {
	if len(e.Name) < 4 {
		return fmt.Errorf("ERROR: Entity needs a unique name (min 3 characters): '%v'", e.Name)
	}

	if _, ok := a.Entities[e.Name]; ok {
		return fmt.Errorf("ERROR: Entity already exists: '%v'", e.Name)
	}

	a.Entities[e.Name] = e
	return nil
}

// AddRelation checks if Entities referenced are existing and adds a new Relation
// to the AST
func (a *Application) AddRelation(rel Relation) error {
	if _, ok := a.Entities[rel.Source]; !ok {
		return fmt.Errorf("ERROR: Source entity does not exists: '%v'", rel.Source)
	}

	if _, ok := a.Entities[rel.Target]; !ok {
		return fmt.Errorf("ERROR: Target entity does not exists: '%v'", rel.Target)
	}

	switch rel.Kind {
	case "onetomany":

	default:
		return fmt.Errorf("ERROR: Missing or unknown relation type: '%v'", rel.Kind)
	}

	name := rel.Source + "_" + rel.Target + "_" + rel.Kind
	if _, ok := a.Relations[name]; ok {
		return fmt.Errorf("ERROR: Relation already exists: '%v'", name)
	}
	a.Relations[name] = rel

	return nil
}

func (a *Application) AddFieldToEntity(entity string, field Field) error {
	// check if entity exists
	if _, ok := a.Entities[entity]; !ok {
		return fmt.Errorf("ERROR: Entity does not exist: '%v'", entity)
	}
	switch field.Kind {
	case "Text":
	case "Password":
	case "Integer":
	case "Number":
	case "Boolean":
	case "Email":
	case "Tel":
	case "Longtext":
	case "Time":
	case "Lookup":

	default:
		return fmt.Errorf("ERROR: Missing or unknown field type: '%v'", field.Kind)
	}

	return nil
}

// SaveToYAML saves the Application definition to a YAML file
func (a Application) SaveToYAML(filepath string) error {
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	if err := a.YAMLWriter(file); err != nil {
		return err
	}

	return nil
}

// YAMLWriter writes the Application struct into a YAML io.Writer stream as []bytes
func (a Application) YAMLWriter(w io.Writer) error {
	enc := yaml.NewEncoder(w)
	defer enc.Close()

	if err := enc.Encode(a); err != nil {
		return fmt.Errorf("stream cannot be encoded into YAML: %v ", err)
	}
	return nil
}

// LoadFromYAML loads the Application definition from a YAML file
func (a *Application) LoadFromYAML(filepath string) error {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return err
	}

	if err := a.YAMLReader(file); err != nil {
		return err
	}

	return nil
}

// YAMLReader reads in the YAML bytes from an io.Reader and converts into
// Application struct
func (a *Application) YAMLReader(r io.Reader) error {
	dec := yaml.NewDecoder(r)
	if err := dec.Decode(a); err != nil {
		return fmt.Errorf("YAML stream cannot be decoded: %v ", err)
	}
	return nil
}