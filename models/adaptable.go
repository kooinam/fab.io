package models

type Adaptable interface {
	NewQuery(*Collection) Queryable
	RegisterCollection(*Collection)
	Collections() []*Collection
}
