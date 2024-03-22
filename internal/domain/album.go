package domain

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}

type AlbumRepository interface {
	SaveAlbum(album *Album) error
	DeleteAlbum(albumId string) error
	GetAlbums() ([]*Album, error)
	GetAlbumById(albumId string) (*Album, error)
	UpdateAlbum(albumId string, album *Album) error
}
