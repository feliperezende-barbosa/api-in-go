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

	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty:   true,
	}

	ctx := context.TODO()
	clientOption := options.Client().ApplyURI(uri).SetBSONOptions(bsonOpts)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	client.Database(dbName).Collection("albums")

	return client
}
