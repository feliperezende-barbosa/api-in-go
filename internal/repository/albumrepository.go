package repository

import "github.com/feliperezende-barbosa/api-in-go/internal/domain"

type AlbumRepo struct {
	handler DBHandler
}

func NewAlbumRepo(d DBHandler) *AlbumRepo {
	return &AlbumRepo{d}
}

// Save implements domain.AlbumRepository.
func (a *AlbumRepo) SaveAlbum(album *domain.Album) error {
	err := a.handler.Save(album)
	if err != nil {
		return err
	}
	return nil
}

func (a *AlbumRepo) DeleteAlbum(albumId string) error {
	err := a.handler.Delete(albumId)
	if err != nil {
		return err
	}
	return nil
}

func (a *AlbumRepo) GetAlbums() ([]*domain.Album, error) {
	albums, err := a.handler.GetAll()
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (a *AlbumRepo) GetAlbumById(albumId string) (*domain.Album, error) {
	albums, err := a.handler.GetById(albumId)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (a *AlbumRepo) UpdateAlbum(albumId string, album *domain.Album) error {
	err := a.handler.Update(albumId, album)
	if err != nil {
		return err
	}
	return nil
}
