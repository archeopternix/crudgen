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
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// Application holds all information and configuration for the AST and consists
// of Entitites, Relations and the Configuration necessary for template generation
type Application struct {
	Name      string              `yaml:"name"`
	Entities  map[string]Entity   `yaml:"entities"`  //Entity
	Relations map[string]Relation `yaml:"relations"` //Relation
	Config    struct {
		PackageName       string `yaml:"packagename"` // just the repo name without the full path
		DateFormat        string `yaml:"dateformat"`
		TimeFormat        string `yaml:"timeformat"`
		CurrencySymbol    string `yaml:"currency_symbol"`
		DecimalSeparator  string `yaml:"decimal_separator"`
		ThousandSeparator string `yaml:"thousand_separator"`
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

// NewFromYAMLFile creates an new Application instance from a YAML file
func NewFromYAMLFile(filepath string) (*Application, error) {
	app := new(Application)
	app.Entities = make(map[string]Entity)
	app.Relations = make(map[string]Relation)

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	if err := app.YAMLReader(file); err != nil {
		return nil, err
	}

	return app, nil
}

// TimeStamp needed for file generation. Will be added in the header of each file
// to track the creation date and time of each file
func (a Application) TimeStamp() string {
	return time.Now().Format(a.Config.DateFormat + " " + a.Config.TimeFormat)
}

// AddEntity adds an new entity to the AST and checks if Entity with this name already exists
// or name is too short
func (a *Application) AddEntity(e Entity) error {
	if e.Kind == "" {
		e.Kind = "default"
	}
	e.Name = strings.ToLower(e.Name)
	e.Kind = strings.ToLower(e.Kind)

	if err := EntityCheckForErrors(e); err != nil {
		return err
	}
	if _, ok := a.Entities[e.Name]; ok {
		return fmt.Errorf("Entity already exists: '%v'", e.Name)
	}

	a.Entities[e.Name] = e
	return nil
}

// AddRelation checks if Entities referenced are existing and adds a new Relation
// to the AST
func (a *Application) AddRelation(rel Relation) error {
	rel.Parent = strings.ToLower(rel.Parent)
	rel.Child = strings.ToLower(rel.Child)
	rel.Kind = strings.ToLower(rel.Kind)

	if _, ok := a.Entities[rel.Parent]; !ok {
		return fmt.Errorf("ERROR: Parent entity does not exists: '%v'", rel.Parent)
	}

	if _, ok := a.Entities[rel.Child]; !ok {
		return fmt.Errorf("ERROR: Target entity does not exists: '%v'", rel.Child)
	}

	switch rel.Kind {
	case "onetomany":

	default:
		return fmt.Errorf("ERROR: Missing or unknown relation type: '%v'", rel.Kind)
	}

	name := rel.Parent + "_" + rel.Child + "_" + rel.Kind
	if _, ok := a.Relations[name]; ok {
		return fmt.Errorf("ERROR: Relation already exists: '%v'", name)
	}
	a.Relations[name] = rel

	return nil
}

// FieldCheckForErrors checks for errors in field definition
func FieldCheckForErrors(f Field) error {

	if f.IsLabel && (!f.Required) {
		return fmt.Errorf("Only required fields can be labels")
	}

	if f.Length < 1 {
		f.Length = 1
	}

	switch f.Kind {
	case "text":
	case "password":
	case "integer":
		if f.Max < f.Min {
			return fmt.Errorf("Max value '%v' must be higher than '%v'", f.Max, f.Min)
		}
	case "number":
		if f.IsLabel {
			return fmt.Errorf("Number cannot be a 'label'")
		}
	case "boolean":
		if f.IsLabel {
			return fmt.Errorf("Boolean cannot be a 'label'")
		}
	case "email":
	case "tel":
		if f.IsLabel {
			return fmt.Errorf("Phone number (tel) cannot be a 'label'")
		}
	case "longtext":
	case "time":
		return fmt.Errorf("Not implemented")
	case "lookup":
		if f.IsLabel {
			return fmt.Errorf("Number cannot be a 'label'")
		}
	default:
		return fmt.Errorf("Missing or unknown field type: '%v'", f.Kind)
	}
	return nil
}

// AddFieldToEntity adds fields to entities and performs some sanity checks
func (a *Application) AddFieldToEntity(entity string, field Field) error {
	entity = strings.ToLower(entity)
	field.Name = strings.ToLower(field.Name)
	field.Object = strings.ToLower(field.Object)
	field.Kind = strings.ToLower(field.Kind)

	// check if entity exists
	if _, ok := a.Entities[entity]; !ok {
		return fmt.Errorf("Entity does not exist: '%v'", entity)
	}

	for _, val := range a.Entities[entity].Fields {
		if (val.Name == field.Name) && (field.Name != "ID") {
			// abort ID field creation without error
			return fmt.Errorf("Field '%v' already exists in entity '%v'", field.Name, entity)
		}
	}

	if err := FieldCheckForErrors(field); err != nil {
		return err
	}

	e := a.Entities[entity]
	e.Fields = append(e.Fields, field)
	a.Entities[entity] = e

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

/* deprecated
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
*/

// YAMLReader reads in the YAML bytes from an io.Reader and converts into
// Application struct
func (a *Application) YAMLReader(r io.Reader) error {
	dec := yaml.NewDecoder(r)
	if err := dec.Decode(a); err != nil {
		return fmt.Errorf("YAML stream cannot be decoded: %v ", err)
	}

	if err := a.parseDependencies(); err != nil {
		return err
	}
	return nil
}

// parseDependencies parse all entities for lookup fields, add unique ID field
// and parse relations between entities and therefore adds dedicated fields for
// parent/child relations and scans for lookups and parent-child relationships
// and therefore creates necessary additional entities (e.g. lookup entities)
// or add additional fields (e.g. Id field for every entity)
func (a *Application) parseDependencies() error {
	for key, entity := range a.Entities {
		for i, field := range entity.Fields {

			// search for lookup fields
			if field.Kind == "lookup" {
				// if entity exists and is not a lookup throw error
				if e, ok := a.Entities[strings.ToLower(field.Name)]; ok {
					if e.Kind != "lookup" {
						return fmt.Errorf("Entity with name '%s' could not be overwritten with 'lookup'", e.Name)
					}
				} else {
					// create new Entity of kind lookup
					a.AddEntity(Entity{
						Name: field.Name,
						Kind: "lookup",
						Fields: []Field{
							{Name: "text", Required: true, Kind: "text", IsLabel: true},
							{Name: "order", Kind: "integer"},
						},
					})
				}
				entity := a.Entities[key]
				entity.Fields[i].Object = entity.Fields[i].Name
				entity.Fields[i].Name = entity.Fields[i].Name + "ID"
				a.Entities[key] = entity
			}
		}
	}

	for key, _ := range a.Entities {
		// add ID field
		a.AddFieldToEntity(key, Field{
			Name:     "ID",
			Kind:     "integer",
			Required: true,
		})
	}

	// add fields for relationships between entities
	for _, relation := range a.Relations {
		if relation.Kind == "onetomany" {
			// add child field
			a.AddFieldToEntity(strings.ToLower(relation.Child), Field{
				Name:   relation.Parent + "ID",
				Kind:   "child",
				Object: relation.Parent,
			})
			// add parent field
			a.AddFieldToEntity(strings.ToLower(relation.Parent), Field{
				Name:   relation.Child,
				Kind:   "parent",
				Object: relation.Child,
			})
		}
	}

	return nil
}
