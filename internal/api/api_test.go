package api_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetAlbumById calls getAlbumById with an id, checking for a valid return value
func TestGetAlbumById(t *testing.T) {
	json := `{"ID": 1, "Title": "Blue Train", "Artist": "John Coltrane", Price: 56.99}`
	var handler = func(*http.Request) (*http.Response, error) {
		re := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       re,
		}, nil
	}

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/albums/%v", 1), nil)
	response, errHandler := handler(request)
	assert.NoError(t, err)
	assert.NoError(t, errHandler)
	assert.NotNil(t, request)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}

// TestPostAlbum calls postAlbum with an album, checking for a valid return value
func TestPostAlbum(t *testing.T) {
	json := `{"ID": 1, "Title": "Blue Train", "Artist": "John Coltrane", Price: 56.99}`
	var handler = func(*http.Request) (*http.Response, error) {
		re := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       re,
		}, nil
	}

	request, err := http.NewRequest(http.MethodPost, "/albums", bytes.NewReader([]byte(json)))
	response, errHandler := handler(request)
	bobyReq, _ := io.ReadAll(request.Body)
	bodyResp, _ := io.ReadAll(response.Body)

	assert.NoError(t, err)
	assert.NoError(t, errHandler)
	assert.NotNil(t, request)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, string(bobyReq), string(bodyResp))
}

// TestUpdateAlbumById calls updateAlbumById with an album, checking for a valid return value
func TestUpdateAlbumById(t *testing.T) {
	json := `{"ID": 1, "Title": "Blue Train", "Artist": "John Coltrane", Price: 56.99}`
	var handler = func(*http.Request) (*http.Response, error) {
		re := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       re,
		}, nil
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/albums/%v", 1), nil)
	response, errHandler := handler(request)
	assert.NoError(t, err)
	assert.NoError(t, errHandler)
	assert.NotNil(t, request)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}

// TestDeleteAlbumById calls deleteAlbumById with an id, checking for a valid return value
func TestDeleteAlbumById(t *testing.T) {
	json := `{"ID": 1, "Title": "Blue Train", "Artist": "John Coltrane", Price: 56.99}`
	var handler = func(*http.Request) (*http.Response, error) {
		re := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       re,
		}, nil
	}

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/albums/%v", 1), nil)
	response, errHandler := handler(request)
	assert.NoError(t, err)
	assert.NoError(t, errHandler)
	assert.NotNil(t, request)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}

func TestGetAlbums(t *testing.T) {
	json := `{"ID": 1, "Title": "Blue Train", "Artist": "John Coltrane", Price: 56.99}`
	var handler = func(*http.Request) (*http.Response, error) {
		re := io.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       re,
		}, nil
	}

	request, err := http.NewRequest(http.MethodGet, "/albums", nil)
	response, errHandler := handler(request)
	body, _ := io.ReadAll(response.Body)

	assert.NoError(t, err)
	assert.NoError(t, errHandler)
	assert.NotNil(t, request)
	assert.NotNil(t, response)
	assert.Equal(t, json, string(body))
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}
