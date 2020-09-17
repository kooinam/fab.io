package simplerecords

import (
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
)

type Query struct {
	collection *models.Collection
}

func makeQuery(collection *models.Collection) *Query {
	query := &Query{
		collection: collection,
	}

	return query
}

// Where used to query collection
func (query *Query) Where(filters helpers.H) models.Queryable {
	panic("simplerecords does not supports Where")
}

// Count used to count records in collection with matching criterion
func (query *Query) Count() *models.CountResult {
	result := models.MakeCountResult()
	result.Set(int64(query.collection.List().Count()), nil)

	return result
}

// ToList used to iterate record in collection with matching criterion
func (query *Query) ToList() *models.ListResults {
	results := models.MakeListResults()
	results.Set(query.collection.List(), nil)

	return results
}

// Each used to iterate record in collection with matching criterion
func (query *Query) Each(handler func(models.Modellable, error) bool) error {
	panic("simplerecords does not supports Each")
}

// First used to return first record in collection with matching criterion
func (query *Query) First() *models.SingleResult {
	panic("simplerecords does not supports First")
}

// FirstOrCreate used to return first record in collection with matching criterion, create one and return if not found
func (query *Query) FirstOrCreate(attributes helpers.H) *models.SingleResult {
	panic("simplerecords does not supports FirstOrCreate")
}

// Find use to find record by id
func (query *Query) Find(id string) *models.SingleResult {
	result := models.MakeSingleResult()
	found := query.collection.List().FindByID(id)

	result.Set(found, nil, found == nil)

	return result
}

// DestroyAll used to destroy all records in collection with matching criterion
func (query *Query) DestroyAll() error {
	panic("simplerecords does not supports DestroyAll")
}
