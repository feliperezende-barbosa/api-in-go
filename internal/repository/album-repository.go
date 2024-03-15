package repository

import "github.com/feliperezende-barbosa/api-in-go/internal/domain"

type AlbumRepo struct {
	handler DbHandler
}

func NewAlbum(d DbHandler) AlbumRepo {
	return AlbumRepo{d}
}

func (a AlbumRepo) SaveAlbum(album domain.Album) error {
	err := a.handler.Save(album)
	if err != nil {
		return err
	}
	return nil
}
