package songService

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nedokyrill/online_library/internal/models/song"
	"github.com/nedokyrill/online_library/internal/repository"
	"io"
	"log"
	"net/http"
	"os"
)

type SongService struct {
	songRepo repository.SongRepository
}

func NewSongService(songRepo repository.SongRepository) *SongService {
	return &SongService{songRepo: songRepo}
}

// @BasePath /api/v1

// @Summary Create song
// @Description Создание песни
// @Tags Song
// @Accept  json
// @Produce  json
// @Param body body song.Song  true  "Данные для создания песни"
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /song/create [post]
func (s *SongService) CreateSong(c *gin.Context) {
	var payload song.Song

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate
	Validator := validator.New()
	if err := Validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error: no mandatory fields": errors})
		return
	}

	if s.songRepo.IsSongExists(payload.Group, payload.Song) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Song of this group already exists"})
		return
	}

	//request to other api
	apiRequest := os.Getenv("OTHER_API") + fmt.Sprintf("?group=%s&song=%s", payload.Group, payload.Song)
	resp, err := http.Get(apiRequest)
	if err != nil {
		log.Println("Error making API request: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch song details from external API"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status: %d", resp.StatusCode)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("API request failed with status: %d", resp.StatusCode)})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	var songDetail song.SongDetail
	if err := json.Unmarshal(body, &songDetail); err != nil {
		log.Println("Error unmarshalling response body: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return
	}

	if songDetail.ReleaseDate != "" || songDetail.Text != "" || songDetail.Link != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no mandatory fields"})
		return
	} else {
		payload.ReleaseDate = songDetail.ReleaseDate
		payload.Text = songDetail.Text
		payload.Link = songDetail.Link
	}

	//query to db
	if err := s.songRepo.CreateSong(payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: Song was not added: ": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message: ": "Song added successfully"})
}
