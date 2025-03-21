package songRepository

import (
	"github.com/nedokyrill/online_library/internal/models/song"
	"log"
)

func (rep *SongRepositoryImpl) UpdateSong(id int64, payload song.Song) error {
	_, err := rep.db.Exec(`UPDATE songs 
	          SET group_name = $1, song_name = $2, release_date = $3, text = $4, link = $5
	          WHERE id = $6`, payload.Group, payload.Song, payload.ReleaseDate, payload.Text, payload.Link, id)
	if err != nil {
		log.Println("Error updating song:", err)
		return err
	}

	log.Println("Song updated successfully")
	return nil
}
