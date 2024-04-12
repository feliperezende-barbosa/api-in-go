package api

import (
	"net/http"

	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"github.com/feliperezende-barbosa/api-in-go/internal/metric"
	"github.com/gin-gonic/gin"
)

type AlbumApi struct {
	albumRepository domain.AlbumRepository
	// telemetryApi    telemetry.Promet
}

func NewAlbumApi(albumRepo domain.AlbumRepository) *AlbumApi {
	return &AlbumApi{albumRepo}
}

func (a *AlbumApi) GetAlbums(c *gin.Context) {
	appMetric := metric.NewHTTP(c.Request.URL.RawPath, c.Request.Method)
	appMetric.Started()

	albums, err := a.albumRepository.GetAlbums()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to fetch albums"})
		return
	}

	c.JSON(http.StatusOK, albums)

	appMetric.Finished()
	appMetric.StatusCode = "200"
}

func (a *AlbumApi) PostAlbums(c *gin.Context) {
	var newAlbum *domain.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := a.albumRepository.SaveAlbum(newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add album"})
		return
	}

	c.JSON(http.StatusCreated, newAlbum)
}

func (a *AlbumApi) GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	album, err := a.albumRepository.GetAlbumById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (a *AlbumApi) UpdateAlbumById(c *gin.Context) {
	id := c.Param("id")

	album := domain.Album{}

	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := a.albumRepository.UpdateAlbum(id, &album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album updated"})
}

func (a *AlbumApi) DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	err := a.albumRepository.DeleteAlbum(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully!"})
}
