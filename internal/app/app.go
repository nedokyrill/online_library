package app

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Run() {
	if err := godotenv.Load(); err != nil { //в будущем можно вытаскивать из докера
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error with connection to database: ", err)
	}
	log.Println("Connected to database successfully")

}
