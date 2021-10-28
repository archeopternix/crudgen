package ast

import (
	"bytes"

	"testing"
)

const res = `name: TestApp
entities: {}
relations: []
config:
  packagename: github.com/archeopternix
`

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
