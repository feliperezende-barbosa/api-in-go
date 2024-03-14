package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Client *mongo.Client
	Db     *mongo.Database
	Albums *mongo.Collection
}

func (c *Mongodb) Conn(uri string, dbName string) {
	client := c.Client
	ctx := context.TODO()
	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	c.Client = client
	c.Db = client.Database(dbName)
	c.Albums = client.Database(dbName).Collection("albums")
}
