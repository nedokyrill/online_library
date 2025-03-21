package services

import "github.com/gin-gonic/gin"

type SongServiceInterface interface {
	GetSongById(c *gin.Context)
	GetAllSongs(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	CreateSong(c *gin.Context)
}
