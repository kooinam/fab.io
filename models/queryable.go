package models

import "github.com/kooinam/fabio/helpers"

// Queryable is the interface for all query adapter implementing query's functionalities
type Queryable interface {
	Where(filter helpers.H) Queryable
	Count() (int64, error)
	Each(func(Modellable, error) bool) error
	First() (Modellable, error)
	FirstOrCreate(helpers.H) (Modellable, error)
	Find(string) (Modellable, error)
}
