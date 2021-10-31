// Package cmd is the command line interface to the CRUD Package generator

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
package cmd

import (
	"crudgen/ast"
	"crudgen/internal"
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

type Environment struct {
	Instance string `yaml:"instance"` // production, development, testing
	Database string `yaml:"database"` // postgres,mysql...
	Host     string `yaml:"host"`     // localhost or IP address
	Port     int    `yaml:"port"`     //5432
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func runModuleCreation(files []string) error {
	if len(files) < 1 {
		return fmt.Errorf("no files to parse are specified: %v", files)
	}
	gen := internal.NewGenerator()

	for _, f := range files {
		if err := gen.ModuleFromYAML(viper.GetString("module-path") + f); err != nil {
			return err
		}
	}

	a, err := ast.NewFromYAMLFile(viper.GetString("cfgpath") + definitionfile)
	if err != nil {
		return err
	}

	gen.Worker = ast.NewGeneratorWorker(a)

	if err := gen.GenerateAll(); err != nil {
		return err
	}

	// initialize or refresh go.mod
	prg := "go"
	arg1 := "mod"
	arg2 := "tidy"
	if internal.FileExist("go.mod") == nil {
		arg2 = "init"
	}

	cmd := exec.Command(prg, arg1, arg2)
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("go.mod (re-)initialized")

	return nil
}

// addField adds an new field to the YAML config file
func addField(f ast.Field) error {

	a, err := ast.NewFromYAMLFile(viper.GetString("cfgpath") + definitionfile)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	if err := a.AddFieldToEntity(entity, f); err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	if err := a.SaveToYAML(viper.GetString("cfgpath") + definitionfile); err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	fmt.Printf("New %v field '%v' added to entity '%v'\n", f.Kind, f.Name, entity)
	return nil
}
