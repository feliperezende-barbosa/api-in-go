package main

import (
	"github.com/feliperezende-barbosa/api-in-go/api"
	"github.com/feliperezende-barbosa/api-in-go/internal/database"
	"github.com/feliperezende-barbosa/api-in-go/internal/repository"
	"github.com/gin-gonic/gin"
)

var (
	dbHanlder database.DBHandler
)

func main() {
	dbHanlder.Conn("mongodb://mongoadmin:mongodbtest@localhost:27017", "test_db")

	r := setupRouter()
	r.Run("localhost:8080")
}

func getAlbumApi() api.AlbumApi {
	albumRepo := repository.NewAlbum(dbHanlder)
	albumApi := api.NewAlbumApi(albumRepo)
	return *albumApi
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	controller := getAlbumApi()

	// router.GET("albums", controller.GetAlbums)
	// router.GET("albums/:id", getAlbumById)
	router.POST("albums", controller.PostAlbums)
	// router.PUT("albums/:id", updateAlbumById)
	// router.DELETE("albums/:id", deleteAlbumById)

	return router
}
