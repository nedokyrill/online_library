package songRepository

import (
	"log"
	"strings"
)

func (rep *SongRepositoryImpl) IsSongExists(group string, songName string) bool {
	normalizedGroupName := strings.ToLower(strings.ReplaceAll(group, " ", ""))
	normalizedSong := strings.ToLower(strings.ReplaceAll(songName, " ", ""))

	var count int
	err := rep.db.QueryRow(`SELECT COUNT(*) FROM songs WHERE LOWER(REPLACE(group_name, ' ', '')) = $1 AND LOWER(REPLACE(song_name, ' ', '')) = $2`, normalizedGroupName, normalizedSong).Scan(&count)
	if err != nil {
		log.Printf("Ошибка при проверке существования песни: %v", err)
		return false
	}

	return count > 0
}
