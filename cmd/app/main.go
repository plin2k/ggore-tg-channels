package main

import (
	"log"
	"os"

	"github.com/plin2k/ggore-tg-channels/internal/models"

	telegram_channel "github.com/plin2k/ggore-tg-channels/internal/app/telegram-channel"

	"github.com/joho/godotenv"
	"github.com/plin2k/ggore-tg-channels/internal/repository"

	"github.com/plin2k/ggore-tg-channels/internal/services"
)

const (
	dbFilePath = "./data/content.db"
	envFile    = ".env"
)

func main() {
	log.Println("Loading app")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalln("Env Init", err)
	}

	cfg := &models.Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
	}

	db, err := repository.NewSQLiteDB(dbFilePath)
	if err != nil {
		log.Fatalf("db initialize error : %s", err.Error())
	}
	defer db.Close()

	repos := repository.NewRepository(db)

	srv, err := services.NewService(cfg)
	if err != nil {
		log.Fatalf("services initialize error : %s", err.Error())
	}

	log.Println("Services loaded")

	log.Println(telegram_channel.Start(srv, repos))
}
