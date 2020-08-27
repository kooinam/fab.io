package models

import "github.com/kooinam/fabio/helpers"

// Queryable is the interface for all query adapter implementing query's functionalities
type Queryable interface {
	Where(filter helpers.H)
	Count() (int64, error)
}
