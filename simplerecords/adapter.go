package simplerecords

import (
	"github.com/kooinam/fabio/models"
)

// Adapter is adapter for simeplerecords
type Adapter struct {
	collections []*models.Collection
	counter     int
}

// MakeAdapter used to instantiate simplerecord's adapter
func MakeAdapter() *Adapter {
	adapter := &Adapter{
		collections: []*models.Collection{},
	}

	return adapter
}

// NewQuery used to generate query
func (adapter *Adapter) NewQuery(collection *models.Collection) models.Queryable {
	panic("simplerecords does not support query")
}

// RegisterCollection used to register collection with adapter
func (adapter *Adapter) RegisterCollection(collection *models.Collection) {
	adapter.collections = append(adapter.collections, collection)
}

// Collections used to retrieve adapter's registered collections
func (adapter *Adapter) Collections() []*models.Collection {
	return adapter.collections
}

func (adapter *Adapter) incrcounter() int {
	adapter.counter++

	return adapter.counter
}
