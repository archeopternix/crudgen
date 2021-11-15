package ast

import (
	"testing"
)

func TestFieldCheckForErrors(t *testing.T) {

	if err := FieldCheckForErrors(Field{Name: "Textfield", Kind: "text"}); err != nil {
		t.Error("Field has to be created")
	}

	if err := FieldCheckForErrors(Field{Name: "Textfield", Kind: "password"}); err != nil {
		t.Error("Field has to be created")
	}
	if err := FieldCheckForErrors(Field{Name: "Textfield", Kind: "email"}); err != nil {
		t.Error("Field has to be created")
	}
	if err := FieldCheckForErrors(Field{Name: "Textfield", Kind: "longtext"}); err != nil {
		t.Error("Field has to be created")
	}
	if err := FieldCheckForErrors(Field{Name: "Textfield", Kind: "tel"}); err != nil {
		t.Error("Field has to be created")
	}

	if err := FieldCheckForErrors(Field{Name: "Textfield", Kind: "text", IsLabel: true}); err == nil {
		t.Error("IsLabel without required must fail")
	}

	if err := FieldCheckForErrors(Field{Name: "1", Kind: "text"}); err == nil {
		t.Error("length of name < 1 must fail")
	}

	if err := FieldCheckForErrors(Field{Name: "A!$§6", Kind: "text"}); err == nil {
		t.Error("name must contain only charcters and number")
	}

	if err := FieldCheckForErrors(Field{Name: "counter", Kind: "integer", Min: 15, Max: 2}); err == nil {
		t.Error("max < 1 min fail")
	}

	if err := FieldCheckForErrors(Field{Name: "counter", Kind: "number", IsLabel: true}); err == nil {
		t.Error("not a label")
	}

	if err := FieldCheckForErrors(Field{Name: "counter", Kind: "boolean", IsLabel: true}); err == nil {
		t.Error("not a label")
	}

	if err := FieldCheckForErrors(Field{Name: "counter", Kind: "tel", IsLabel: true}); err == nil {
		t.Error("not a label")
	}

	if err := FieldCheckForErrors(Field{Name: "counter", Kind: "lookup", IsLabel: true}); err == nil {
		t.Error("not a label")
	}

	if err := FieldCheckForErrors(Field{Name: "counter", Kind: "time"}); err == nil {
		t.Error("time not implemented yet")

	}

	if err := FieldCheckForErrors(Field{Name: "alpha", Kind: "dinosaur"}); err == nil {
		t.Error("unknown field type")
	}

}
