package database

import (
	"context"
	"log"

	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHandler struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var (
	ctx = context.TODO()
)

func (m *MongoHandler) Conn(uri string, dbName string) {
	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	m.Client = client
	m.Db = client.Database(dbName)
}

func (m *MongoHandler) Save(album domain.Album) error {
	collection := m.Db.Collection("albums")
	_, err := collection.InsertOne(ctx, album)
	if err != nil {
		return err
	}
	return nil
}
