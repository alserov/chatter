package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func MustConnect(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic("failed to ping db: " + err.Error())
	}

	return client
}
