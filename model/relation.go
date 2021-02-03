// Package model implements the model definitions for the scuffold project
package model

// RelationType is the type of relation
// currently supported types are : 	One2many, Many2many
type RelationType int

// Relation type definitions
const (
	One2many = iota
	Many2many
)

func (r RelationType) String() string {
	return [...]string{"One2many", "Many2many"}[r]
}

// Relation defines the relation between 2 entity object
type Relation struct {
	Name       string       `yaml:"name"`
	ParentName string       `yaml:"parent"`
	ChildName  string       `yaml:"child"`
	Kind       RelationType `yaml:"kind"` // "one2many", "many2many"
}
