package mongorecords

import (
	"github.com/kooinam/fab.io/helpers"
	"github.com/kooinam/fab.io/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Query is a wrapper for querying mongo
type Query struct {
	collection *models.Collection
	filters    helpers.H
	sorts      helpers.H
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
func (query *Query) Count() *models.CountResult {
	result := models.MakeCountResult()

	adapter := query.collection.Adapter().(*Adapter)
	collection := adapter.getCollection(query.collection.Name())
	ctx := adapter.getTimeoutContext()

	count, err := collection.CountDocuments(ctx, query.filters)

	result.Set(count, err)

	return result
}

// ToList used to iterate record in collection with matching criterion
func (query *Query) ToList() *models.ListResults {
	results := models.MakeListResults()
	list := models.MakeList()

	adapter := query.collection.Adapter().(*Adapter)
	collection := adapter.getCollection(query.collection.Name())
	ctx := adapter.getTimeoutContext()
	cursor, err := collection.Find(ctx, query.filters)

	for cursor.Next(ctx) {
		item := query.collection.New(helpers.H{})

		err := cursor.Decode(item)

		if err != nil {
			break
		} else {
			list.Add(item)
		}
	}

	results.Set(list, err)

	return results
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
func (query *Query) First() *models.SingleResult {
	result := models.MakeSingleResult()

	adapter := query.collection.Adapter().(*Adapter)
	collection := adapter.getCollection(query.collection.Name())
	ctx := adapter.getTimeoutContext()
	item := query.collection.New(helpers.H{})

	err := collection.FindOne(ctx, query.filters).Decode(item)

	result.Set(item, err, query.haveNotFound(err))

	return result
}

// FirstOrCreate used to return first record in collection with matching criterion, create one and return if not found
func (query *Query) FirstOrCreate(attributes helpers.H) *models.SingleResult {
	result := query.First()

	if result.StatusNotFound() {
		// not found, create one
		result = query.collection.Create(helpers.Merge(query.filters, attributes))
	}

	return result
}

// Find use to find record by id
func (query *Query) Find(id string) *models.SingleResult {
	result := models.MakeSingleResult()

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		result.Set(nil, err, false)
	}

	adapter := query.collection.Adapter().(*Adapter)
	collection := adapter.getCollection(query.collection.Name())
	ctx := adapter.getTimeoutContext()
	item := query.collection.New(helpers.H{})

	err = collection.FindOne(ctx, helpers.Merge(query.filters, helpers.H{
		"_id": oid,
	})).Decode(item)

	result.Set(item, err, query.haveNotFound(err))

	return result
}

// DestroyAll used to destroy all records in collection with matching criterion
func (query *Query) DestroyAll() error {
	var err error

	// TODO

	return err
}

// Sort used to sort collection
func (query *Query) Sort(field string, asc bool) models.Queryable {

	return query
}

func (query *Query) haveNotFound(err error) bool {
	if err != nil && err.Error() == "mongo: no documents in result" {
		return true
	}

	return false
}
