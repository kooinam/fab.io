package mongorecords

import (
	"context"
	"time"

	"github.com/kooinam/fab.io/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Adapter is adapter for mongo
type Adapter struct {
	client       *mongo.Client
	uri          string
	databaseName string
	collections  map[string]*models.Collection
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
		collections:  make(map[string]*models.Collection),
	}

	return adapter, err
}

// NewQuery used to generate query
func (adapter *Adapter) NewQuery(collection *models.Collection) models.Queryable {
	query := makeQuery(collection)

	return query
}

// RegisterCollection used to register collection with adapter
func (adapter *Adapter) RegisterCollection(collectionName string, collection *models.Collection) {
	adapter.collections[collectionName] = collection
}

// Collection used to retrieve registered collection
func (adapter *Adapter) Collection(collectionName string) *models.Collection {
	return adapter.collections[collectionName]
}

// Collections used to retrieve registered collections
func (adapter *Adapter) Collections() []*models.Collection {
	collections := []*models.Collection{}

	for _, collection := range adapter.collections {
		collections = append(collections, collection)
	}

	return collections
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
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	return err
}
