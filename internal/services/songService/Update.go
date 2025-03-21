package songService

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/nedokyrill/online_library/internal/models/song"
	Utils "github.com/nedokyrill/online_library/pkg/utils"
	"net/http"
)

// @BasePath /api/v1

// @Summary Update song
// @Description Обновление данных песни
// @Tags Song
// @Accept  json
// @Produce  json
// @Param id path int true "ID песни"
// @Param body body song.Song  true  "Данные для обновления песни"
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /song/{id}/update [patch]
func (s *SongService) Update(c *gin.Context) {
	id := Utils.GetIDFromContext(c)
	var payload song.Song

	//verification of existence
	existingSong, err := s.songRepo.GetSongById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with getting song": err.Error()})
		return
	}

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// copy data with ignore nil-fields
	err = copier.CopyWithOption(existingSong, payload, copier.Option{IgnoreEmpty: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//query to db
	if err := s.songRepo.UpdateSong(id, *existingSong); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error with updating song": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "song updated successfully"})

}
