package models

type Adaptable interface {
	NewQuery(*Collection) Queryable
	RegisterCollection(string, *Collection)
	Collection(string) *Collection
	Collections() []*Collection
}
