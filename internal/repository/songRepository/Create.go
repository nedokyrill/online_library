package songRepository

import (
	"database/sql"
	"github.com/nedokyrill/online_library/internal/models/song"
	"log"
)

type SongRepositoryImpl struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepositoryImpl {
	return &SongRepositoryImpl{db}
}

func (rep *SongRepositoryImpl) CreateSong(payload song.Song) error {
	_, err := rep.db.Exec(
		"INSERT INTO songs (group_name, song_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5)",
		payload.Group,
		payload.Song,
		payload.ReleaseDate,
		payload.Text,
		payload.Link,
	)

	if err != nil {
		log.Println("Error inserting song: ", err)
		return err
	}

	log.Println("Song inserted successfully")
	return nil
}
