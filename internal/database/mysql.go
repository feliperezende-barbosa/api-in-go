package database

import (
	"database/sql"
	"log"

	"github.com/feliperezende-barbosa/api-in-go/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type MySqlHandler struct {
	db *sql.DB
}

func (ms *MySqlHandler) Conn() {
	mysqlCfg := mySqlConfig()
	database, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	ms.db = database
}

func mySqlConfig() *mysql.Config {
	cfg := mysql.Config{
		User:   "",
		Passwd: "",
		Net:    "tcp",
		Addr:   "",
		DBName: "album",
	}
	return &cfg
}

// Delete implements repository.DBHandler.
func (ms MySqlHandler) Delete(albumId string) error {
	row := ms.db.QueryRow("DELETE FROM album WHERE id = ?", albumId)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

// GetAll implements repository.DBHandler.
func (ms MySqlHandler) GetAll() ([]*domain.Album, error) {
	var albums []*domain.Album

	rows, err := ms.db.Query("SELECT * FROM album")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb *domain.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return albums, nil
}

// GetById implements repository.DBHandler.
func (ms MySqlHandler) GetById(albumId string) (*domain.Album, error) {
	row := ms.db.QueryRow("SELECT * FROM album WHERE id = ?", albumId)
	if row.Err() != nil {
		return nil, row.Err()
	}

	album := domain.Album{}
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &album, nil
}

// Save implements repository.DBHandler.
func (ms MySqlHandler) Save(album domain.Album) error {
	_, err := ms.db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return err
	}
	return nil
}

// Update implements repository.DBHandler.
func (ms MySqlHandler) Update(albumId string, album domain.Album) error {
	_, err := ms.db.Exec("UPDATE album SET title = ?, artist = ?, price = ?) WHERE id = ?", album.Title, album.Artist, album.Price, albumId)
	if err != nil {
		return err
	}
	return nil
}
