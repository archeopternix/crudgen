package model

import (
	"os"
	. "scuffold/model"

	"bytes"
	"time"

	"gopkg.in/yaml.v3"
)

// Datastore interface has to be implemented for loading and storing the 'Application'
// data to a persistent store
type YAMLDatastore struct {
	File string
}

func NewYAMLDatastore(filename string) (*YAMLDatastore, error) {
	/*var file *os.File
	var err error
	file, err = os.Open(filename) // For read access.
	if err != nil {
		return nil, err // File cannot be opened
	}
	defer file.Close() */
	ds := new(YAMLDatastore)
	ds.File = filename
	return ds, nil
}

func (ds *YAMLDatastore) RetrieveAllData() (*Application, error) {
	var b bytes.Buffer
	var app *Application

	file, err := os.Open(ds.File) // For read access.
	if err != nil {
		return nil, err // File cannot be opened
	}
	defer file.Close()

	_, err = file.Read(b.Bytes())
	if err != nil {
		return nil, err // YAML cannot be read from file
	}
	app = new(Application)
	err = yaml.Unmarshal(b.Bytes(), app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (ds *YAMLDatastore) SaveAllData(app *Application) error {
	file, err := os.Create(ds.File) // For write access.
	if err != nil {
		return err // File cannot be opened
	}
	defer file.Close()
	app.Timestamp = time.Now()
	b, er := yaml.Marshal(*app)
	if er != nil {
		return er // Application connot be converted into YAML
	}
	_, err = file.Write(b)
	if err != nil {
		return err // YAML cannot be written into file
	}
	return nil
}
