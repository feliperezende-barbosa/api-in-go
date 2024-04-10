package repository_test

import (
	"testing"

	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var albums = []domain.Album{
	{ID: uuid.New(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: uuid.New(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: uuid.New(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type MockAlbumRepository struct {
	mock.Mock
}

func (mock *MockAlbumRepository) SaveAlbum(album *domain.Album) error {
	args := mock.Called()
	return args.Error(1)
}

func (mock *MockAlbumRepository) DeleteAlbum(albumId string) error {
	args := mock.Called()
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
	return args.Error(1)
}

// TestGetAlbums calls getAlbums, checking for a valid return value
func TestGetAlbums(t *testing.T) {
	mockRepo := new(MockAlbumRepository)
	album := domain.Album{ID: uuid.New(), Title: "Test", Artist: "Test", Price: 0}
	mockRepo.On("GetAlbums").Return([]*domain.Album{&album}, nil)

	testService := domain.AlbumRepository(mockRepo)
	result, _ := testService.GetAlbums()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, result, []*domain.Album{&album})
}

func TestGetAlbumById(t *testing.T) {
	mockRepo := new(MockAlbumRepository)
	mockRepo.On("GetAlbumById").Return(&albums[1], nil)

	testService := domain.AlbumRepository(mockRepo)
	result, _ := testService.GetAlbumById(albums[1].ID.String())

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result)
	assert.Equal(t, result, &albums[1])

}

func TestSaveAlbum(t *testing.T) {
	mockRepo := new(MockAlbumRepository)
	mockRepo.On("SaveAlbum").Return(&albums[1], nil)

	testService := domain.AlbumRepository(mockRepo)
	err := testService.SaveAlbum(&albums[1])

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestDeleteAlbum(t *testing.T) {
	mockRepo := new(MockAlbumRepository)
	mockRepo.On("DeleteAlbum").Return(&albums[1], nil)

	testService := domain.AlbumRepository(mockRepo)
	err := testService.DeleteAlbum(albums[1].ID.String())

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestUpdateAlbum(t *testing.T) {
	mockRepo := new(MockAlbumRepository)
	mockRepo.On("UpdateAlbum").Return(&albums[1], nil)

	testService := domain.AlbumRepository(mockRepo)
	err := testService.UpdateAlbum(albums[1].ID.String(), &albums[1])

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}
