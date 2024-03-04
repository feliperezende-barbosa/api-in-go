package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Conn(uri string, dbName string) (mongo.Client, error) {

	ctx := context.TODO()
	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
		return *client, err
	}

	client.Database(dbName)

	return *client, err

}
