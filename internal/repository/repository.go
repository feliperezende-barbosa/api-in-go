package repository

import "github.com/feliperezende-barbosa/api-in-go/internal/domain"

type DbHandler interface {
	GetAll() ([]*domain.Album, error)
	GetById(album domain.Album) error
	Save(album domain.Album) error
	Update(album domain.Album) error
	Delete(album domain.Album) error
}
