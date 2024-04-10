package main

import (
	"github.com/feliperezende-barbosa/api-in-go/internal/api"
	"github.com/feliperezende-barbosa/api-in-go/internal/database"
	"github.com/feliperezende-barbosa/api-in-go/internal/repository"
	"github.com/gin-gonic/gin"
)

var (
	dbHanlder *database.MongoHandler
	// mySqlHandler *database.MySqlHandler
)

func main() {
	dbHanlder.Conn("mongodb://mongoadmin:mongodbtest@localhost:27017", "test_db")
	// mySqlHandler.Conn()

	r := setupRouter()
	r.Run("localhost:8080")
}

func getAlbumApi() *api.AlbumApi {
	albumRepo := repository.NewAlbumRepo(dbHanlder)
	albumApi := api.NewAlbumApi(albumRepo)
	return albumApi
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	controller := getAlbumApi()

	router.GET("albums", controller.GetAlbums)
	router.GET("albums/:id", controller.GetAlbumById)
	router.POST("albums", controller.PostAlbums)
	router.PUT("albums/:id", controller.UpdateAlbumById)
	router.DELETE("albums/:id", controller.DeleteAlbumById)

	return router
}
