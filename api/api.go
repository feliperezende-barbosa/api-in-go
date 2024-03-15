package main

import (
	"net/http"

	"github.com/feliperezende-barbosa/api-in-go/internal/database"
	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	RepoMongo database.MongoHandler
)

func main() {

	RepoMongo.Conn("mongodb://mongoadmin:mongodbtest@localhost:27017", "test_db")

	r := setupRouter()

	r.Run("localhost:8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("albums", getAlbums)
	router.GET("albums/:id", getAlbumById)
	router.POST("albums", postAlbums)
	router.PUT("albums/:id", updateAlbumById)
	router.DELETE("albums/:id", deleteAlbumById)

	return router
}

func getAlbums(c *gin.Context) {
	var albums []*domain.Album

	cursor, err := RepoMongo.Db.Collection("albums").Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch albums"})
		return
	}

	if err = cursor.All(c, &albums); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch albums"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum domain.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	_, err := RepoMongo.Db.Collection("albums").InsertOne(c, newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add album"})
		return
	}

	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	filter := bson.M{"id": id}

	album := domain.Album{}

	res := RepoMongo.Db.Collection("albums").FindOne(c, filter)
	err := res.Decode(&album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

func updateAlbumById(c *gin.Context) {
	id := c.Param("id")

	album := domain.Album{}

	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	filter := bson.M{"id": id}
	update := bson.M{"$set": album}

	_, err := RepoMongo.Db.Collection("albums").UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album updated"})
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	filter := bson.M{"id": id}

	_, err := RepoMongo.Db.Collection("albums").DeleteOne(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully!"})
}
