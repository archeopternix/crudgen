// Package model implements the model definitions for the scuffold project
package model

import (
	"fmt"
)

// EntityType is the type of entity
// currently supported types are : Regular, Lookup
type EntityType int

// Entity types separates between lookup and regular entities
const (
	Regular = iota
	Lookup
)

func (e EntityType) String() string {
	return [...]string{"Regular", "Lookup"}[e]
}

// Entity relates to an 'Object' or struct
type Entity struct {
	Name       string            `yaml:"name"`
	Fields     map[string]*Field `yaml:",flow"`
	EntityType `yaml:"type"`     // 0..Normal, 1..Lookup
}

// AddField add a field to an existing entity. Checks if a field with the same
// name already exists  for this entity
func (e *Entity) AddField(field *Field) error {
	if _, ok := e.Fields[field.Name]; ok {
		return fmt.Errorf("model: entity %s already contais field: %s", e.Name, field.Name)
	}
	e.Fields[field.Name] = field
	return nil

}
