// Package model implements the model definitions for the scuffold project
package model

import (
	"fmt"
	"time"
)

// Settings is the definition of the global attributes
type Settings struct {
	CurrencySymbol    string `yaml:"currency_symbol"`
	DecimalSeparator  string `yaml:"decimal_separator"`
	ThousendSeparator string `yaml:"thousend_separator"`
	TimeFormat        string `yaml:"time_format"`
}

// Timestamp constants (can be adjusted)
//    ANSIC       = "Mon Jan _2 15:04:05 2006"
//    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
//    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
//    RFC822      = "02 Jan 06 15:04 MST"
//    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
//    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
//    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
//    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
//    RFC3339     = "2006-01-02T15:04:05Z07:00"
//    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
//    Kitchen     = "3:04PM"
//    Stamp      = "Jan _2 15:04:05"
//    StampMilli = "Jan _2 15:04:05.000"
//    StampMicro = "Jan _2 15:04:05.000000"
//    StampNano  = "Jan _2 15:04:05.000000000"
type Application struct {
	Timestamp time.Time            `yaml:"timestamp"`
	Globals   Settings             `yaml:"globals"`   // GlobalSettings for application generation
	Entities  map[string]*Entity   `yaml:"entities"`  // Entities holds all entity definitions
	Relations map[string]*Relation `yaml:"relations"` // Relations holds in a map all relations between entities
}

// NewEntity creates a new entity, adds it to the Entities map and
// checks if a entity with the same name already exists
func (app *Application) NewEntity(name string, e EntityType) (*Entity, error) {
	entity := new(Entity)
	entity.EntityType = e
	entity.Name = name
	entity.Fields = make(map[string]*Field)
	if _, ok := app.Entities[name]; ok {
		return nil, fmt.Errorf("model: entity already exists: %s", name)
	}
	app.Entities[name] = entity
	return entity, nil
}

// NewRelation creates a new relation, adds it to the Relations map and
// checks if a relation with the same name already exists
// The name of the entity consists of the type, name of parent and name of child entity
func (app *Application) NewRelation(parent string, child string, r RelationType) (*Relation, error) {
	relation := new(Relation)
	relation.Kind = r
	relation.ParentName = parent
	relation.ChildName = child
	name := r.String() + "_" + parent + "_" + child
	if _, ok := App.Relations[name]; ok {
		return nil, fmt.Errorf("model: relation already exists: %s", name)
	}
	relation.Name = name
	App.Relations[name] = relation
	return relation, nil
}

var App *Application

func init() {
	App = new(Application)
	App.Globals = Settings{CurrencySymbol: "€", DecimalSeparator: ",", ThousendSeparator: ".", TimeFormat: "02.01.2006 15:04:05"}
	App.Relations = make(map[string]*Relation)
	App.Entities = make(map[string]*Entity)
}
