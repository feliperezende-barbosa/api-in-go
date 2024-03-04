package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func conn() {

	ctx := context.TODO()
	clientOption := options.Client().ApplyURI("")

	client, err := mongo.Connect(ctx, clientOption)
	if err == nil {
		return client
	}
	return err
}
