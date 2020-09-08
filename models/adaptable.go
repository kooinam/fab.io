package models

// Adaptable is the interface for all adapters implementing adapter's functionalities
type Adaptable interface {
	NewQuery(*Collection) Queryable
	RegisterCollection(string, *Collection)
	Collection(string) *Collection
	Collections() []*Collection
}
