package main

import (
	"fmt"
	"net/http"

	"github.com/feliperezende-barbosa/api-in-go/internal/database/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
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
	cursor, err := mongodb.Albums.Find(c, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch products"})
		return
	}

	var albums []album
	if err = cursor.All(c, &albums); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch products"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func updateAlbumById(c *gin.Context) {
	id := c.Param("id")

	var updateAlbum album

	if err := c.BindJSON(&updateAlbum); err != nil {
		return
	}

	for index, a := range albums {
		if a.ID == id {
			albums = append(albums[:index], updateAlbum)
			c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Album %v uptated successfully!", a.Title)})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func deleteAlbumById(c *gin.Context) {

	id := c.Param("id")

	for index, a := range albums {
		if a.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted successfully!"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
