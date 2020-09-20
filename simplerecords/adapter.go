package simplerecords

import (
	"github.com/kooinam/fab.io/models"
)

// Adapter is adapter for simeplerecords
type Adapter struct {
	collections map[string]*models.Collection
	counter     int
}

// MakeAdapter used to instantiate simplerecord's adapter
func MakeAdapter() *Adapter {
	adapter := &Adapter{
		collections: make(map[string]*models.Collection),
	}

	return adapter
}

// NewQuery used to generate query
func (adapter *Adapter) NewQuery(collection *models.Collection) models.Queryable {
	query := makeQuery(collection)

	return query
}

// RegisterCollection used to register collection with adapter
func (adapter *Adapter) RegisterCollection(collectionName string, collection *models.Collection) {
	adapter.collections[collectionName] = collection
}

// Collection used to retrieve registered collection
func (adapter *Adapter) Collection(collectionName string) *models.Collection {
	return adapter.collections[collectionName]
}

// Collections used to retrieve registered collections
func (adapter *Adapter) Collections() []*models.Collection {
	collections := []*models.Collection{}

	for _, collection := range adapter.collections {
		collections = append(collections, collection)
	}

	return collections
}

func (adapter *Adapter) incrcounter() int {
	adapter.counter++

	return adapter.counter
}
