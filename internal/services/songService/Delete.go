package songService

import (
	"github.com/gin-gonic/gin"
	Utils "github.com/nedokyrill/online_library/pkg/utils"
	"net/http"
)

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
