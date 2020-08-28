package mongorecords

import (
	"context"
	"time"

	"github.com/kooinam/fabio/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Adapter is adapter for mongo
type Adapter struct {
	client       *mongo.Client
	uri          string
	databaseName string
	collections  []*models.Collection
}

// MakeAdapter used to instantiate mongorecord's adapter
func MakeAdapter(uri string, databaseName string) (*Adapter, error) {
	var err error
	var adapter *Adapter
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	adapter = &Adapter{
		client:       client,
		uri:          uri,
		databaseName: databaseName,
		collections:  []*models.Collection{},
	}

	return adapter, err
}

// NewQuery used to generate query
func (adapter *Adapter) NewQuery(collection *models.Collection) models.Queryable {
	query := makeQuery(collection)

	return query
}

// RegisterCollection used to register collection with adapter
func (adapter *Adapter) RegisterCollection(collection *models.Collection) {
	adapter.collections = append(adapter.collections, collection)
}

// Collections used to retrieve adapter's registered collections
func (adapter *Adapter) Collections() []*models.Collection {
	return adapter.collections
}

func (adapter *Adapter) getCollection(collectionName string) *mongo.Collection {
	collection := adapter.client.Database(adapter.databaseName).Collection(collectionName)

	return collection
}

func (adapter *Adapter) getTimeoutContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return ctx
}

func (adapter *Adapter) save(base *Base) error {
	var err error
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_ = context

	return err
}
