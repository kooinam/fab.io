package mongo

import (
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
)

// Query is a wrapper for querying mongo
type Query struct {
	collection *models.Collection
	filters    helpers.H
	err        error
}

func makeQuery(collection *models.Collection) *Query {
	return &Query{
		collection: collection,
		filters:    helpers.H{},
	}
}

// Where used to query collection
func (query *Query) Where(filters helpers.H) models.Queryable {
	query.filters = filters

	return query
}

// Count used to count records in collection with matching criterion
func (query *Query) Count() (int64, error) {
	adapter := query.collection.Adapter().(*Adapter)

	collection := adapter.getCollection(query.collection.Name())
	ctx := adapter.getTimeoutContext()

	return collection.CountDocuments(ctx, query.filters)
}

// Each used to iterate record in collection with matching criterion
func (query *Query) Each(handler func(models.Modellable, error) bool) error {
	var err error

	adapter := query.collection.Adapter().(*Adapter)

	collection := adapter.getCollection(query.collection.Name())
	ctx := adapter.getTimeoutContext()
	cursor, err := collection.Find(ctx, query.filters)

	for cursor.Next(ctx) {
		var err2 error
		item := query.collection.New(helpers.H{})

		err2 = cursor.Decode(item)

		shouldContinue := handler(item, err2)

		if !shouldContinue {
			break
		}
	}

	return err
}

// First used to return first record in collection with matching criterion
func (query *Query) First() (models.Modellable, error) {
	var err error

	adapter := query.collection.Adapter().(*Adapter)

	collection := adapter.getCollection(query.collection.Name())
	ctx := adapter.getTimeoutContext()
	item := query.collection.New(helpers.H{})

	err = collection.FindOne(ctx, query.filters).Decode(item)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (query *Query) FirstOrCreate(values helpers.H) (models.Modellable, error) {
	err := query.err

	// TODO

	return nil, err
}

func (query *Query) Find(id string) (models.Modellable, error) {
	err := query.err

	// TODO

	return nil, err
}

func (query *Query) DestroyAll() error {
	var err error

	// TODO

	return err
}
