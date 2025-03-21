package songRepository

import (
	"database/sql"
	"github.com/nedokyrill/online_library/internal/models/song"
	"log"
)

func (rep *SongRepositoryImpl) GetSongById(id int64) (*song.Song, error) {
	var Song song.Song
	err := rep.db.QueryRow("SELECT id, group_name, song_name, release_date, text, link FROM songs WHERE id = $1", id).Scan(
		&Song.Id,
		&Song.Group,
		&Song.Song,
		&Song.ReleaseDate,
		&Song.Text,
		&Song.Link)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no songs founded")
			return nil, err
		} else {
			log.Println("error getting song")
			return nil, err
		}
	}
	return &Song, nil
}
