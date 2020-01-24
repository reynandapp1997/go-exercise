package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func Connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://127.0.0.1:27017/")
	clientOptions.SetConnectTimeout(time.Second * 15)
	clientOptions.SetServerSelectionTimeout(time.Second * 15)
	clientOptions.SetSocketTimeout(time.Second * 15)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("crud_golang"), nil
}
func GetContext() context.Context {
	return ctx
}
