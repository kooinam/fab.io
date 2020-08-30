package models

import "github.com/kooinam/fabio/helpers"

// Queryable is the interface for all query adapter implementing query's functionalities
type Queryable interface {
	Where(filter helpers.H) Queryable
	Count() *CountResult
	Each(func(Modellable, error) bool) error
	First() *SingleResult
	FirstOrCreate(helpers.H) *SingleResult
	Find(string) *SingleResult
	All() *ListResults
}
