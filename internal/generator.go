// Package internal consists of the full AST (abstract syntax tree) which reflects
// the object structure consisting of Entities, Fields, Relations..

/*
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
package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Generator for file generation, holds all Modules and TaskWorkers which are
// implemented by specific models
type Generator struct {
	Modules map[string]Module
	Worker  TaskWorker
}

// TaskWorker interface has to be implemented by Application class
type TaskWorker interface {
	Generate(task *Task) error
}

// Module is one independent dedicated functional unit that holds all Tasks
// (activities) to generate a certain part of an application (e.g. HTML view, Entities)
type Module struct {
	path  string
	Name  string `yaml:"name"`
	Tasks []Task
}

// Task is a single task for file generation which could be the copy of file or
// the generation based on template execution.
//
// Currently 2 modes are supported 'copy' or 'template'.
// Appdate = true indicates that the whole Application structure is submitted to
// the template generator. When Filename is set (not nil) the whole Application
// will be send to the template execution. If Filename is empty the generator
// iterates over all Entities and calls the template generator with a single entity.
// Filename is provided without file extension
type Task struct {
	Kind     string   `yaml:"kind"` // currently supported: copy, template
	Source   []string `yaml:"source"`
	Target   string   `yaml:"target"`             // target directory - filename wil be calculated
	Template string   `yaml:"template,omitempty"` // name of the template from template file
	Fileext  string   `yaml:"fileext,omitempty"`  // file extension for the generated file
	Filename string   `yaml:"filename,omitempty"` // when Filename is set (not nil) the whole Application will be send to the template execution
}

// NewGenerator creates a Generator
func NewGenerator() *Generator {
	g := new(Generator)
	g.Modules = make(map[string]Module)

	return g
}

// ModuleFromYAML reads in a 'Module' from an YAML file and adds it to the generator configuration
// In a post processing step source/target filenames and filepaths will be cleaned
func (c *Generator) ModuleFromYAML(filename string) error {
	r, err := os.Open(filename)
	defer r.Close()
	if err != nil {
		return fmt.Errorf("YAML file %v could not be loaded: #%v ", filename, err)
	}

	if m, err := ModuleFromReader(r, filepath.Dir(filename)); err == nil {
		c.Modules[m.Name] = *m
		log.Printf("read in module '%s' from file '%s'", m.Name, filename)
	} else {
		return err
	}

	return nil
}

// ModuleFromReader reads in a 'Module' from an YAML stream and cleans
// source/target filenames and filepaths
func ModuleFromReader(reader io.Reader, path string) (*Module, error) {
	m := new(Module)
	m.path = path

	dec := yaml.NewDecoder(reader)
	if err := dec.Decode(m); err != nil {
		return nil, fmt.Errorf("could not be decoded: #%v ", err)
	}

	var tasks []Task
	for _, t := range m.Tasks {
		tasks = append(tasks, t.CleanPaths(m.path))
	}
	m.Tasks = tasks

	return m, nil
}

// SaveToFile saves the full generator configuration to a YAML file
func (c Generator) saveToFile(filename string) error {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be marshalled: #%v ", filename, err)
	}

	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be saved: #%v ", filename, err)
	}

	log.Printf("saved Generator configuration to file '%s'", filename)
	return nil
}

// GenerateAll calls the GenerateModule function of all Modules
func (c Generator) GenerateAll() error {
	for _, m := range c.Modules {
		for _, t := range m.Tasks {
			if err := c.Worker.Generate(&t); err != nil {
				return err
			}
		}
	}
	return nil
}

// AddTask adds a task to a module an clen the target and source path
func (m *Module) AddTask(t *Task) {
	m.Tasks = append(m.Tasks, t.CleanPaths(m.path))
}

// CleanPaths cleans the target and source path and adds to the sourcepath the
// filepath of the module
// - source: the module path will be added so fields are accessible from root of application
// - target: just clean the target path
func (t *Task) CleanPaths(modulepath string) Task {

	task := *t
	task.Source = nil
	for _, s := range t.Source {
		task.Source = append(task.Source, filepath.Join(modulepath, filepath.Clean(s)))
	}
	task.Target = filepath.Clean(t.Target)
	return task
}
