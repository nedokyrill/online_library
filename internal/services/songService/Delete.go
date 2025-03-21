package songService

import (
	"github.com/gin-gonic/gin"
	Utils "github.com/nedokyrill/online_library/pkg/utils"
	"net/http"
)

// @BasePath /api/v1

// @Summary Delete song
// @Description Удаление песни
// @Tags Song
// @Accept  json
// @Produce  json
// @Param id path int true "ID песни"
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /song/{id}/delete [delete]
func (s *SongService) Delete(c *gin.Context) {
	id := Utils.GetIDFromContext(c)

	//verification of existence
	_, err := s.songRepo.GetSongById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with getting song": err.Error()})
		return
	}

	//query to db
	if err := s.songRepo.DeleteSong(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: song was not deleted: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Song successfully, song's id: ": id})
}
