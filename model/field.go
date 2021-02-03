// Package model implements the model definitions for the scuffold project
package model

// FieldType is the type of Field
// currently supported types are: String, Integer, Bool, Dropdown, Phone, Email
type FieldType int

func (f FieldType) String() string {
	return [...]string{"String", "Integer", "Bool", "Dropdown", "Phone", "Email"}[f]
}

// Field type definitions
const (
	String = iota
	Integer
	Bool
	Dropdown
	Phone
	Email
)

// Field is each and every single attribute.
// Object is empty except in case type=lookup or child keeps the name of the Object
type Field struct {
	Name      string `yaml:"name"`
	FieldType `json:"fieldtype"`
	Object    string `yaml:"object,omitempty"` // for lookup, child relations - mappingtable for many2many relations
	Maxlength int    `yaml:"maxlength,omitempty"`
	Size      int    `yaml:"size,omitempty"` // for textarea size = cols
	Required  bool   `yaml:"required"`
	Step      int    `yaml:"step,omitempty"`    //for Number fields
	Min       int    `yaml:"min,omitempty"`     //for Number fields
	Max       int    `yaml:"max,omitempty"`     //for Number fields
	Rows      int    `yaml:"rows,omitempty"`    //for textarea
	IsLabel   bool   `yaml:"islabel,omitempty"` // when true is the shown text for select boxes
}
