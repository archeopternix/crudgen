package ast

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Holds all entitites or one dedicated Object for template generation
type Application struct {
	Name      string
	Entities  map[string]Entity //Entity
	Relations []string          //Relation
	Config    struct {
		PackageName string
	}
}

func NewApplication(name string) *Application {
	app := new(Application)
	app.Name = name
	app.Entities = make(map[string]Entity)
	return app
}

func (a Application) SaveToYAML(filepath string) error {
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}

	enc := yaml.NewEncoder(file)
	defer enc.Close()
	err = enc.Encode(a)
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) LoadFromYAML(filepath string) error {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be loaded: #%v ", filepath, err)
	}
	err = yaml.Unmarshal(file, a)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be unmarshalled: #%v ", filepath, err)
	}

	return nil
}
