package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Adapter is adapter for mongo
type Adapter struct {
	client       *mongo.Client
	uri          string
	databaseName string
}

func makeAdapter(uri string, databaseName string) (*Adapter, error) {
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
	}

	return adapter, err
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
