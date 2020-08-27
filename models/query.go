package models

import (
	"github.com/kooinam/fabio/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

// Query is a wrapper for querying mongo
type Query struct {
	adapter    *Adapter
	collection *Collection
	filters    bson.M
	err        error
}

// Where used to query collection
func (query *Query) Where(filters bson.M) *Query {
	query.filters = filters

	return query
}

// Count used to count records in collection with matching criterion
func (query *Query) Count() (int64, error) {
	err := query.err

	if err != nil {
		return 0, err
	}

	collection := query.adapter.getCollection(query.collection.name)
	ctx := query.adapter.getTimeoutContext()

	return collection.CountDocuments(ctx, query.filters)
}

// Each used to iterate record in collection with matching criterion
func (query *Query) Each(handler func(Modellable, error)) error {
	err := query.err

	if err != nil {
		return err
	}

	collection := query.adapter.getCollection(query.collection.name)
	ctx := query.adapter.getTimeoutContext()
	cursor, err := collection.Find(ctx, query.filters)

	for cursor.Next(ctx) {
		var err2 error
		item := query.collection.New(helpers.H{})

		err2 = cursor.Decode(item)

		handler(item, err2)
	}

	return err
}

// First used to return first record in collection with matching criterion
func (query *Query) First() (Modellable, error) {
	err := query.err

	collection := query.adapter.getCollection(query.collection.name)
	ctx := query.adapter.getTimeoutContext()
	item := query.collection.New(helpers.H{})

	err = collection.FindOne(ctx, query.filters).Decode(item)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (query *Query) FirstOrCreate(values helpers.H) (Modellable, error) {
	err := query.err

	// TODO

	return nil, err
}

func (query *Query) Find(id string) (Modellable, error) {
	err := query.err

	// TODO

	return nil, err
}

func (query *Query) DestroyAll() error {
	var err error

	// TODO

	return err
}
