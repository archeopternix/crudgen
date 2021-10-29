package internal

import (
	"bytes"
	"fmt"
	"testing"
)

const res = `# Configuration
name: Test
tasks:
# Main
- kind: template
  fileext: .go
  source:
  - main.tmpl
  target: 
  template: main
  filename: main

`

type TestWorker struct {
}

func (tw TestWorker) Generate(task *Task) error {
	return nil
}

func TestGenerator(t *testing.T) {
	g := NewGenerator()
	if g == nil {
		t.Errorf("Generator not created")
	}

	buf := bytes.NewBuffer([]byte(res))
	m, err := ModuleFromReader(buf, "core")
	if err != nil {
		t.Errorf("Module not read: %v", err)
	}

	g.Modules[m.Name] = *m
	g.Worker = TestWorker{}

	if err := g.GenerateAll(); err != nil {
		t.Errorf("GenerateAll throws error: %v", err)
	}

}

func TestModule(t *testing.T) {
	buf := bytes.NewBuffer([]byte(res))

	m, err := ModuleFromReader(buf, "core")
	if err != nil {
		t.Errorf("Module not read: %v", err)
	}

	if m.Name != "Test" {
		t.Errorf("Module 'Test' name not found: %s", m.Name)
	}

	if m.Tasks[0].Source[0] != "core\\main.tmpl" {
		t.Errorf("Task source 'core\\main.tmpl' not found: %s", m.Tasks[0].Source[0])
	}
}
