package database

import (
	"context"

	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHandler struct {
	db *mongo.Database
}

var (
	ctx = context.TODO()
)

func NewMongoDB(uri string, dbName string) *MongoHandler {
	clientOption := *options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, &clientOption)
	if err != nil {
		panic(err)
	}
	return &MongoHandler{
		db: client.Database(dbName),
	}
}

func (m *MongoHandler) Save(album *domain.Album) error {
	collection := m.db.Collection("albums")
	_, err := collection.InsertOne(ctx, album)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements repository.DBHandler.
func (m *MongoHandler) Delete(albumId string) error {
	collection := m.db.Collection("albums")
	filter := bson.M{"id": albumId}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements repository.DBHandler.
func (m *MongoHandler) GetAll() ([]*domain.Album, error) {
	collection := m.db.Collection("albums")
	var albums []*domain.Album

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &albums); err != nil {
		return nil, err
	}
	return albums, nil
}

// GetById implements repository.DBHandler.
func (m *MongoHandler) GetById(albumId string) (*domain.Album, error) {
	collection := m.db.Collection("albums")
	filter := bson.M{"id": albumId}

	album := domain.Album{}
	res := collection.FindOne(ctx, filter)
	err := res.Decode(&album)
	if err != nil {
		return nil, err
	}

	return &album, nil
}

// Update implements repository.DBHandler.
func (m *MongoHandler) Update(albumId string, album *domain.Album) error {
	collection := m.db.Collection("albums")

	filter := bson.M{"id": albumId}
	update := bson.M{"$set": album}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
