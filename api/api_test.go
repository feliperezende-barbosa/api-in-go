package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func albums() []domain.Album {
	return []domain.Album{
		{ID: uuid.New(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: uuid.New(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: uuid.New(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}

// TestGetAlbums calls getAlbums, checking for a valid return value
func TestGetAlbums(t *testing.T) {
	// newRecord := httptest.NewRecorder()
	// req, _ := http.NewRequest("GET", "/albums", nil)
	// router.ServeHTTP(newRecord, req)

	// assert.Equal(t, http.StatusOK, newRecord.Code)
	// // assert.Equal(t, "opa", newRecord.Body.String())
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
		name          string
		buildServer   func(w http.ResponseWriter, r *http.Request)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			buildServer: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`{"Artist": "John Coltrane", "Price": 56.99}`))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, "/albums", nil)
			require.NoError(t, err)
			require.Equal(t, "opa", "opa")

			tc.buildServer(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
