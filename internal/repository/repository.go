package repository

import "github.com/nedokyrill/online_library/internal/models/song"

type SongRepository interface {
	GetSongById(id int64) (*song.Song, error)
	GetAllSongs(filters song.Song, limit, offset int) ([]song.Song, error)
	UpdateSong(id int64, payload song.Song) error
	DeleteSong(id int64) error
	CreateSong(payload song.Song) error
	IsSongExists(groupName string, songName string) bool
}
