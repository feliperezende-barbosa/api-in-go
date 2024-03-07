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

	// bsonOpts := &options.BSONOptions{
	// 	UseJSONStructTags: true,
	// 	NilSliceAsEmpty:   true,
	// }

	// uri = "mongodb://localhost:27017"
	// dbName = "test_db"

	ctx := context.TODO()
	clientOption := options.Client().ApplyURI("mongodb://<user>:<pass>@localhost:27017") //.SetBSONOptions(bsonOpts)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	Albums = client.Database("test_db").Collection("albums")

	return client
}
