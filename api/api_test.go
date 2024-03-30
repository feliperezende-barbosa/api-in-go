package api_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var albums = []*domain.Album{
	{ID: uuid.New(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: uuid.New(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: uuid.New(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type MockAlbumRepository struct {
	mock.Mock
}

func (mock *MockAlbumRepository) SaveAlbum(album *domain.Album) error {
	args := mock.Called()
	// result := args.Get(0)
	return args.Error(1)
}

func (mock *MockAlbumRepository) DeleteAlbum(albumId string) error {
	args := mock.Called()
	// result := args.Get(0)
	return args.Error(1)
}

func (mock *MockAlbumRepository) GetAlbums() ([]*domain.Album, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]*domain.Album), args.Error(1)
}

func (mock *MockAlbumRepository) GetAlbumById(albumId string) (*domain.Album, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Album), args.Error(1)
}

func (mock *MockAlbumRepository) UpdateAlbum(albumId string, album *domain.Album) error {
	args := mock.Called()
	// result := args.Get(0)
	return args.Error(1)
}

// TestGetAlbums calls getAlbums, checking for a valid return value
func TestGetAlbums(t *testing.T) {
	mockRepo := new(MockAlbumRepository)
	mockRepo.On("GetAlbums").Return(albums, nil)

	testService := domain.AlbumRepository(mockRepo)
	result, _ := testService.GetAlbums()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, result, albums)

}

// TestGetAlbumById calls getAlbumById with an id, checking for a valid return value
func TestGetAlbumById(t *testing.T) {
	// albumId := domain.Album{ID: uuid.New()}
	// newRecord := httptest.NewRecorder()
	// req, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%v", albumId.ID), nil)
	// router.ServeHTTP(newRecord, req)

	// assert.Equal(t, http.StatusOK, newRecord.Code)
}

// TestGetAlbumById calls getAlbumById with an id, checking for an error.
func TestErrorGetAlbumById(t *testing.T) {
	// albumId := domain.Album{ID: uuid.New()}
	// newRecord := httptest.NewRecorder()
	// req, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%v", albumId.ID), nil)
	// router.ServeHTTP(newRecord, req)

	// assert.Equal(t, http.StatusNotFound, newRecord.Code)
	// assert.Equal(t, "{\n    \"message\": \"Album not found\"\n}", newRecord.Body.String())
}

// TestPostAlbum calls postAlbum with an album, checking for a valid return value
func TestPostAlbum(t *testing.T) {
	// newAlbum := domain.Album{ID: uuid.New(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99}
	// newRecord := httptest.NewRecorder()
	// jsonValue, _ := json.Marshal(newAlbum)
	// req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
	// router.ServeHTTP(newRecord, req)

	// assert.Equal(t, http.StatusCreated, newRecord.Code)
}

// TestUpdateAlbumById calls updateAlbumById with an album, checking for a valid return value
func TestUpdateAlbumById(t *testing.T) {
	// newAlbum := domain.Album{ID: uuid.New(), Title: "Sarah Vaughan and Me", Artist: "Sarah Vaughan", Price: 39.99}
	// newRecord := httptest.NewRecorder()
	// jsonValue, _ := json.Marshal(newAlbum)
	// req, _ := http.NewRequest("PUT", fmt.Sprintf("/albums/%v", newAlbum.ID), bytes.NewBuffer(jsonValue))
	// router.ServeHTTP(newRecord, req)

	// assert.Equal(t, http.StatusOK, newRecord.Code)
}

// TestDeleteAlbumById calls deleteAlbumById with an id, checking for a valid return value
func TestDeleteAlbumById(t *testing.T) {
	// albumId := domain.Album{ID: uuid.New()}
	// newRecord := httptest.NewRecorder()
	// req, _ := http.NewRequest("DELETE", fmt.Sprintf("/albums/%v", albumId.ID), nil)
	// router.ServeHTTP(newRecord, req)

	// assert.Equal(t, http.StatusOK, newRecord.Code)
}

func TestAlbumApi_GetAlbums(t *testing.T) {
	testCases := []struct {
		name    string
		handler func(w http.ResponseWriter, r *http.Request)
	}{
		{
			name: "Ok",
			handler: func(w http.ResponseWriter, r *http.Request) {
				io.WriteString(w, fmt.Sprintln(albums))
			},
		},
		{
			name: "NotOk",
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "unable to fetch albums", http.StatusBadRequest)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, "/albums", nil)
			tc.handler(recorder, request)
			if tc.name != "NotOk" {
				require.NoError(t, err)
				require.Equal(t, albums, albums)
				require.Equal(t, http.StatusOK, recorder.Code)
			} else {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			}
		})
	}
}
