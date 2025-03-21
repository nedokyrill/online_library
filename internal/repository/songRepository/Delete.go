package songRepository

import "log"

func (rep *SongRepositoryImpl) DeleteSong(id int64) error {
	_, err := rep.db.Exec("DELETE FROM songs WHERE id = $1", id)

	if err != nil {
		log.Println("error with deleting song: ", err)
		return err
	}

	log.Println("song deleted successfully")
	return nil
}
