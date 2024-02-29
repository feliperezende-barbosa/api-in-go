package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetAlbums calls getAlbums, checking for a valid return value
func TestGetAlbums(t *testing.T) {
	router := setupRouter()

	newRecord := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album", nil)
	router.ServeHTTP(newRecord, req)

	assert.Equal(t, http.StatusOK, newRecord.Code)
}

// TestGetAlbumById calls getAlbumById with an id, checking for a valid return value
func TestGetAlbumById(t *testing.T) {
	router := setupRouter()

	albumId := album{ID: "1"}

	newRecord := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/album/%v", albumId.ID), nil)
	router.ServeHTTP(newRecord, req)

	assert.Equal(t, http.StatusOK, newRecord.Code)
}

// TestGetAlbumById calls getAlbumById with an id, checking for an error
func TestErrorGetAlbumById(t *testing.T) {
	router := setupRouter()

	albumId := album{ID: "0"}

	newRecord := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/album/%v", albumId.ID), nil)
	router.ServeHTTP(newRecord, req)

	assert.Equal(t, http.StatusNotFound, newRecord.Code)
	assert.Equal(t, "{\n    \"message\": \"Album not found\"\n}", newRecord.Body.String())
}

// TestPostAlbum calls postAlbum with an album, checking for a valid return value
func TestPostAlbum(t *testing.T) {
	router := setupRouter()

	newAlbum := album{ID: "4", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99}

	newRecord := httptest.NewRecorder()
	jsonValue, _ := json.Marshal(newAlbum)
	req, _ := http.NewRequest("POST", "/album", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(newRecord, req)

	assert.Equal(t, http.StatusCreated, newRecord.Code)
}

// TestUpdateAlbumById calls updateAlbumById with an album, checking for a valid return value
func TestUpdateAlbumById(t *testing.T) {
	router := setupRouter()

	newAlbum := album{ID: "4", Title: "Sarah Vaughan and Me", Artist: "Sarah Vaughan", Price: 39.99}

	newRecord := httptest.NewRecorder()
	jsonValue, _ := json.Marshal(newAlbum)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/album/%v", newAlbum.ID), bytes.NewBuffer(jsonValue))
	router.ServeHTTP(newRecord, req)

	assert.Equal(t, http.StatusOK, newRecord.Code)
}

// TestDeleteAlbumById calls deleteAlbumById with an id, checking for a valid return value
func TestDeleteAlbumById(t *testing.T) {
	router := setupRouter()

	albumId := album{ID: "3"}

	newRecord := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/album/%v", albumId.ID), nil)
	router.ServeHTTP(newRecord, req)

	assert.Equal(t, http.StatusOK, newRecord.Code)
}
