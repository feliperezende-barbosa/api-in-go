package repository

import "github.com/feliperezende-barbosa/api-in-go/internal/domain"

type DBHandler interface {
	GetAll() ([]*domain.Album, error)
	GetById(albumId string) (*domain.Album, error)
	Save(album *domain.Album) error
	Update(albumId string, album *domain.Album) error
	Delete(albumId string) error
}
