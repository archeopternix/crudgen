// Package ast consists of the full AST (abstract syntax tree) which reflects
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
package ast

import (
	"crudgen/internal"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
)

// GeneratorWorker generates files based on definition in Task
type GeneratorWorker struct {
	app Application
}

func NewGeneratorWorker(a *Application) *GeneratorWorker {
	g := new(GeneratorWorker)
	g.app = *a
	return g
}

// GenerateModule generates a 'Module' based on the Generator configuration.
// Currently implemented Modules are:
//
// kind: copy:
// - source: contains all the source files that will be copied 1:1
// - target: is the path where all source files will be copied into
//
// kind: template:
// - source: contains all the template files that will be used
// - target: is the path where all source files will be copied into
// - template: name of the primary template used for generation {{define "kinds"}}
// - fileext: is the extension of the generated files
// - filename (optional): when ilename is set (not nil) the whole Application will be send to the template execution
func (gw GeneratorWorker) Generate(task *internal.Task) error {
	// pluralize and singularize functions for templates
	pl := pluralize.NewClient()
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		"lowercase": strings.ToLower, "singular": pl.Singular, "title": strings.Title, "plural": pl.Plural, "inc": func(counter int) int { return counter + 1 },
	}

	// check or create path
	if err := internal.CheckMkdir(task.Target); err != nil {
		_, ok := err.(*internal.DirectoryExistError)
		if ok {
			log.Printf("directory '%s' already exists\n", task.Target)
		} else {
			return err
		}
	} else {
		log.Printf("directory '%s' created\n", task.Target)
	}

	switch task.Kind {
	case "copy":
		// copying all files from .Source to .Target
		for _, src := range task.Source {
			path := filepath.Join(task.Target, filepath.Base(src))
			if err := internal.CopyFile(src, path); err != nil {
				_, ok := err.(*internal.FileExistError)
				if ok {
					log.Println(err)
				} else {
					return err
				}
			} else {
				log.Printf("file '%s' created\n", path)
			}
		}
	case "template":
		// Create a template, add the function map, and parse the text.
		tmpl, err := template.New(task.Template).Funcs(funcMap).ParseFiles(task.Source...)
		if err != nil {
			log.Fatalf("parsing: %s", err)
		}

		if len(task.Filename) > 0 {
			file := filepath.Join(task.Target, strings.ToLower(task.Filename)+task.Fileext)
			writer, err := os.Create(file)
			if err != nil {
				return fmt.Errorf("template generator %v", err)
			}
			defer writer.Close()
			if err := tmpl.ExecuteTemplate(writer, task.Template, gw.app); err != nil {
				return fmt.Errorf("templategenerator %v", err)
			}
			log.Printf("template '%s' written to file '%s'\n", task.Template, file)

		} else {
			for _, entity := range gw.app.Entities {
				file := filepath.Join(task.Target, strings.ToLower(entity.Name)) + task.Fileext
				writer, err := os.Create(file)
				if err != nil {
					return fmt.Errorf("template generator %v", err)
				}
				defer writer.Close()
				entityStruct := struct {
					Entity
					AppName   string
					TimeStamp string
				}{
					entity,
					gw.app.Name,
					gw.app.TimeStamp(),
				}
				if err := tmpl.ExecuteTemplate(writer, task.Template, entityStruct); err != nil {
					return fmt.Errorf("templategenerator %v", err)
				}
				log.Printf("template '%s' for entity '%s' written to file '%s'\n", task.Template, entity.Name, file)
			}
		}
	default:
		return fmt.Errorf("unknown generator operation '%s'", task.Kind)
	}

	return nil
}
