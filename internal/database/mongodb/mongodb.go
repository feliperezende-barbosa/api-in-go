package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Albums *mongo.Collection
)

func Conn(uri string, dbName string) *mongo.Client {

	ctx := context.TODO()
	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	Albums = client.Database(dbName).Collection("albums")

	return client
}
