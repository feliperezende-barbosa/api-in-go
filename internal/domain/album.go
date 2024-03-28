package domain

import "github.com/google/uuid"

type Album struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
}

type AlbumRepository interface {
	SaveAlbum(album *Album) error
	DeleteAlbum(albumId string) error
	GetAlbums() ([]*Album, error)
	GetAlbumById(albumId string) (*Album, error)
	UpdateAlbum(albumId string, album *Album) error
}
