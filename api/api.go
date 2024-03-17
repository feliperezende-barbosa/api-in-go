package api

import (
	"net/http"

	"github.com/feliperezende-barbosa/api-in-go/internal/database"
	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type AlbumApi struct {
	albumRepository domain.AlbumRepository
}

var (
	RepoMongo database.MongoHandler
)

func NewAlbumApi(albumRepo domain.AlbumRepository) *AlbumApi {
	return &AlbumApi{albumRepo}
}

func GetAlbums(c *gin.Context) {
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

func (a *AlbumApi) PostAlbums(c *gin.Context) {
	var newAlbum domain.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := a.albumRepository.Save(newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add album"})
		return
	}

	c.JSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
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

func UpdateAlbumById(c *gin.Context) {
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

func DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	filter := bson.M{"id": id}

	_, err := RepoMongo.Db.Collection("albums").DeleteOne(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully!"})
}
