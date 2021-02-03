// Repository provides the abstract interface to the Application which stores all entity,
// field and relation data
package model

import (
	. "scuffold/model"
)

// Datastore interface has to be implemented for loading and storing the 'Application'
// data to a persistent store
type Datastore interface {
	RetrieveAllData() (*Application, error)
	SaveAllData(app *Application) error
}
