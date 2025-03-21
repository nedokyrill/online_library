package songRepository

import (
	"fmt"
	"github.com/nedokyrill/online_library/internal/models/song"
	"log"
)

func (rep *SongRepositoryImpl) GetAllSongs(filters song.Song, limit, offset int) ([]song.Song, error) {
	log.Println(filters.Id)
	query, args := buildQuery(filters.Id, filters.Group, filters.Song, filters.ReleaseDate, filters.Text, filters.Link, limit, offset)
	log.Println(query, args)
	rows, err := rep.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var songs []song.Song
	for rows.Next() {
		var s song.Song
		if err := rows.Scan(&s.Id, &s.Group, &s.Song, &s.ReleaseDate, &s.Text, &s.Link); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		songs = append(songs, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return songs, nil
}

func buildQuery(id int, group, song, releaseDate, text, link string, limit, offset int) (string, []interface{}) {
	query := "SELECT * FROM songs WHERE 1=1"
	args := []interface{}{}
	i := 1

	// Добавляем условия в запрос для каждого фильтра
	if id != 0 {
		query += fmt.Sprintf(" AND id = $%d", i)
		args = append(args, id)
		i++
	}
	if group != "" {
		query += fmt.Sprintf(" AND group LIKE = $%d", i)
		args = append(args, "%"+group+"%")
		i++
	}
	if song != "" {
		query += fmt.Sprintf(" AND song LIKE $%d", i)
		args = append(args, "%"+song+"%")
		i++
	}
	if releaseDate != "" {
		query += fmt.Sprintf(" AND releaseDate = $%d", i)
		args = append(args, releaseDate)
		i++
	}
	if text != "" {
		query += fmt.Sprintf(" AND text LIKE $%d", i)
		args = append(args, "%"+text+"%")
		i++
	}
	if link != "" {
		query += fmt.Sprintf(" AND link LIKE $%d", i)
		args = append(args, "%"+link+"%")
		i++
	}

	// Добавляем пагинацию
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	return query, args
}
