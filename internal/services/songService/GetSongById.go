package songService

import (
	"github.com/gin-gonic/gin"
	Utils "github.com/nedokyrill/online_library/pkg/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// @BasePath /api/v1

// @Summary Get song
// @Description Просмотреть песни по id
// @Tags Song
// @Accept  json
// @Produce  json
// @Param id path int true "ID песни"
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /song/{id}/ [get]
func (s *SongService) GetSongById(c *gin.Context) {
	id := Utils.GetIDFromContext(c)
	log.Println("1")
	// get limit n offset from context
	limitStr := c.DefaultQuery("limit", "1")
	offsetStr := c.DefaultQuery("offset", "0")

	log.Println("2")
	song, err := s.songRepo.GetSongById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// validate limit n offset
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
		return
	}

	// split text
	verses := Utils.SplitVerses(song.Text)

	// pagination
	if offset >= len(verses) {
		offset = limit
	}
	if offset+limit > len(verses) {
		limit = len(verses) - offset
	}

	paginatedText := strings.Join(verses[offset:limit], "\n")

	c.JSON(http.StatusOK, gin.H{
		"id":          song.Id,
		"group":       song.Group,
		"song":        song.Song,
		"releaseDate": song.ReleaseDate,
		"verses":      paginatedText,
		"link":        song.Link,
	})
}
