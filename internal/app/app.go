package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nedokyrill/online_library/internal/handlers/songHandler"
	"github.com/nedokyrill/online_library/internal/repository/songRepository"
	"github.com/nedokyrill/online_library/internal/services/songService"
	Utils "github.com/nedokyrill/online_library/pkg/utils"

	"log"
	"os"
)

func Run() {
	if err := godotenv.Load(); err != nil { //в будущем можно вытаскивать из докера
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error with connection to database: ", err)
	}
	log.Println("Connected to database successfully", os.Getenv("DATABASE_URL"))

	//repo
	songRepo := songRepository.NewSongRepository(db)

	//service
	songSrv := songService.NewSongService(songRepo)

	//handler
	songHand := songHandler.NewSongHandler(songSrv)

	//default route
	router := gin.Default()
	api := router.Group("/api/v1/")

	////swagger
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//docs.SwaggerInfo.BasePath = "/api/v1"

	//REGISTER_ROUTES
	songHand.RegisterRoutes(api)

	server := Utils.NewServer(os.Getenv("ADDR"), router)
	log.Println("Server success running on port: ", os.Getenv("ADDR"))
	Utils.Start(server)
}
