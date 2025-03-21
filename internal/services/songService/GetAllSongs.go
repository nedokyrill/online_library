package songService

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/online_library/internal/models/song"
	"log"
	"net/http"
	"strconv"
)

// @BasePath /api/v1

// @Summary Get All Songs
// @Description Просмотреть все песни
// @Tags Song
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Param id query int false "Фильтр по Id"
// @Param group query string false "Фильтр по Group"
// @Param song query string false "Фильтр по Song"
// @Param release_date query string false "Фильтр по ReleaseDate"
// @Param text query string false "Фильтр по Text"
// @Param link query string false "Фильтр по Link"
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /song/search-all [get]
func (s *SongService) GetAllSongs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

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

	var filters song.Song

	if id := c.Query("id"); id != "" {
		id, err := strconv.Atoi(id)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id value"})
			return
		}
		filters.Id = id
	} else {
		filters.Id = 0
	}
	filters.Group = c.Query("group")
	filters.Song = c.Query("song")
	filters.ReleaseDate = c.Query("release_date")
	filters.Text = c.Query("text")
	filters.Link = c.Query("link")

	log.Println(filters.Id)
	songs, err := s.songRepo.GetAllSongs(filters, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"songs": songs})
}
