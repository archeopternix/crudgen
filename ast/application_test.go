package ast

import (
	"bytes"

	"testing"
)

const res = `name: TestApp
entities: {}
relations: {}
config:
  packagename: github.com/archeopternix
  dateformat: ""
  timeformat: ""
  currency_symbol: ""
  decimal_separator: ""
  thousand_separator: ""
`

func TestTimeStamp(t *testing.T) {
	a := NewApplication("TestApp")
	a.Config.DateFormat = "02.01.2006"
	a.Config.TimeFormat = "15:04:05.000"
	res := a.TimeStamp()
	if len(res) != 23 {
		t.Errorf("Timestamp expected to be in format: %v result was: '%s", a.Config.DateFormat+" "+a.Config.TimeFormat, res)
	} else {
		t.Logf("Timestamp format: '%s' result '%s", a.Config.DateFormat+" "+a.Config.TimeFormat, res)
	}

}

func TestYAMLWriter(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewApplication("TestApp")
	a.Config.PackageName = "github.com/archeopternix"
	if err := a.YAMLWriter(buf); err != nil {
		t.Errorf("Conversion to YAML: %v", err)
	}
	if buf.String() != res {
		t.Errorf("Result does not match. Result:\n %v\nExpected:\n %v", buf, res)
	} else {
		t.Logf("YAML matches expected result.")
	}
}

func TestYAMLReader(t *testing.T) {
	b := []byte(res)
	buf := bytes.NewBuffer(b)
	var a Application

	if err := a.YAMLReader(buf); err != nil {
		t.Errorf("Conversion from YAML to Struct: %v", err)
	}

	if a.Name != "TestApp" {
		t.Errorf("Result does not match. Name result: %v expected: TestApp", a.Name)
	} else {
		t.Logf("YAML matches expected result.")
	}
}

func TestAddEntity(t *testing.T) {
	a := NewApplication("TestApp")

	// Add a new entity
	if err := a.AddEntity(Entity{Name: "Alpha"}); err != nil {
		t.Errorf("Creation of entity failed: %v", err)
	}
	if len(a.Entities) < 1 {
		t.Error("Adding Entitiy failed")
	}

	// Add the same entity again (should fail)
	if err := a.AddEntity(Entity{Name: "Alpha"}); err == nil {
		t.Errorf("Creation of duplicate entity must fail")
	}

	// Add entity with too short name
	if err := a.AddEntity(Entity{Name: "A1"}); err == nil {
		t.Error("Creation of entity with too short name must fail")
	}

	// Add entity with too short name
	if err := a.AddEntity(Entity{Name: "A1", Kind: "dinosaur"}); err == nil {
		t.Error("Creation of entity with unknown kind must fail")
	}
}

func TestAddRelation(t *testing.T) {
	a := NewApplication("TestApp")

	// Add a new entity
	if err := a.AddEntity(Entity{Name: "Alpha"}); err != nil {
		t.Errorf("Creation of entity 'Alpha' failed: %v", err)
	}
	// Add a new entity
	if err := a.AddEntity(Entity{Name: "Beta"}); err != nil {
		t.Errorf("Creation of entity 'Beta' failed: %v", err)
	}

	// Add a relation with unknown Entities
	if err := a.AddRelation(Relation{Parent: "Beta", Child: "Zeta", Kind: "onetomany"}); err == nil {
		t.Errorf("Unknown targt entity has to fail")
	}

	// Add a relation with missing relation type
	if err := a.AddRelation(Relation{Parent: "Beta", Child: "Alpha"}); err == nil {
		t.Errorf("Entity with unknown relation type has to fail")
	}

	// Add a relation
	if err := a.AddRelation(Relation{Parent: "Beta", Child: "Alpha", Kind: "onetomany"}); err != nil {
		t.Errorf("Creation of ralation failed: %v", err)
	}
	if len(a.Relations) < 1 {
		t.Error("Adding Relation failed")
	}

	// Add the same relation again
	if err := a.AddRelation(Relation{Parent: "Beta", Child: "Alpha", Kind: "onetomany"}); err == nil {
		t.Errorf("Duplicate entities must not created")
	}
}

func TestAddTextField(t *testing.T) {
	a := NewApplication("TestApp")

	// Add a new entity
	if err := a.AddEntity(Entity{Name: "Alpha"}); err != nil {
		t.Errorf("Creation of entity 'Alpha'  failed: %v", err)
	}

	// unknow entity
	if err := a.AddFieldToEntity("Gamma", Field{Name: "Textfield"}); err == nil {
		t.Errorf("Creation of Field 'Textfield' not failed: %v", err)
	}

	// unknown/missing type
	if err := a.AddFieldToEntity("Alpha", Field{Name: "Textfield"}); err == nil {
		t.Errorf("Creation of Field 'Textfield' not failed: %v", err)
	}

	// successful
	if err := a.AddFieldToEntity("Alpha", Field{Name: "Textfield", Kind: "text"}); err != nil {
		t.Errorf("Creation of Field 'Textfield' failed: %v", err)
	}
}

func TestAddIDField(t *testing.T) {
	a := NewApplication("TestApp")

	// Add a new entity
	if err := a.AddEntity(Entity{Name: "Alpha"}); err != nil {
		t.Errorf("Creation of entity 'Alpha'  failed: %v", err)
	}

	// successful
	if err := a.AddFieldToEntity("Alpha", Field{Name: "ID", Kind: "integer"}); err != nil {
		t.Errorf("Creation of Field 'Integer' failed: %v", err)
	}
	// successful
	if err := a.AddFieldToEntity("Alpha", Field{Name: "ID", Kind: "integer"}); err != nil {
		t.Errorf("ID field must not created but should not throw any error %v", err)
	}
}
