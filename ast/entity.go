package ast

// Entity relates to an 'Object' or struct
type Entity struct {
	Name string `yaml:"name"`
	// Fields []Field `yaml:"fields"`
	Kind string `yaml:"type,omitempty"` // 0..Normal, 1..Lookup 2..Many2Many
}
