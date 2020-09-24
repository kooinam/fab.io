package simplerecords

import (
	"fmt"

	"github.com/kooinam/fab.io/helpers"
	"github.com/kooinam/fab.io/models"
)

type Query struct {
	collection *models.Collection
	filters    helpers.H
}

func makeQuery(collection *models.Collection) *Query {
	query := &Query{
		collection: collection,
	}

	return query
}

// Where used to query collection
func (query *Query) Where(filters helpers.H) models.Queryable {
	query.filters = filters

	return query
}

// Count used to count records in collection with matching criterion
func (query *Query) Count() *models.CountResult {
	result := models.MakeCountResult()
	result.Set(int64(query.ToList().List().Count()), nil)

	return result
}

// ToList used to iterate record in collection with matching criterion
func (query *Query) ToList() *models.ListResults {
	results := models.MakeListResults()
	list := query.collection.List()

	newList := list.FindAll(func(item models.Modellable) bool {
		matched := true

		for key, value := range query.filters {
			fieldValue := helpers.GetFieldValueByName(item, key)

			if fieldValue != value {
				matched = false
			}
		}

		return matched
	})

	results.Set(newList, nil)

	return results
}

// Each used to iterate record in collection with matching criterion
func (query *Query) Each(handler func(models.Modellable, error) bool) error {
	var err error

	items := query.collection.List().Items()

	for _, item := range items {
		shouldContinue := handler(item, nil)

		if !shouldContinue {
			break
		}
	}

	return err
}

// First used to return first record in collection with matching criterion
func (query *Query) First() *models.SingleResult {
	result := models.MakeSingleResult()

	list := query.collection.List()

	found := list.Find(func(item models.Modellable) bool {
		matched := true

		for key, value := range query.filters {
			fieldValue := helpers.GetFieldValueByName(item, key)

			if fieldValue != value {
				matched = false
			}
		}

		return matched
	})

	if found != nil {
		result.Set(found, nil, false)
	} else {
		result.Set(found, fmt.Errorf("item not found"), true)
	}

	return result
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
