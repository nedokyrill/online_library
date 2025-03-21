package songHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/online_library/internal/services"
	Utils "github.com/nedokyrill/online_library/pkg/utils"
)

type SongHandler struct {
	service services.SongServiceInterface
}

func NewSongHandler(service services.SongServiceInterface) *SongHandler {
	return &SongHandler{
		service: service,
	}
}

func (h *SongHandler) RegisterRoutes(router *gin.RouterGroup) {
	songHandler := router.Group("/song")
	{
		songHandler.GET("/search-all", h.service.GetAllSongs)
		songHandler.POST("/create", h.service.CreateSong)

		songIdHandler := songHandler.Group("/:id")
		songIdHandler.Use(Utils.ValidateAndExtractIDMiddleware())
		{
			songIdHandler.GET("/", h.service.GetSongById)
			songIdHandler.PATCH("/update", h.service.Update)
			songIdHandler.DELETE("/delete", h.service.Delete)
		}
	}
}
