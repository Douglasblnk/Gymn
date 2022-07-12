package main

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := database.InitDB()

	if err != nil {
		log.Fatal("db", err)
	}

	user := &models.User{
		Name:        "Fulano",
		Email:       "mock@mock.com",
		Password:    "123",
		Is_personal: false,
	}

	if err := db.Select("*").Create(user).Error; err != nil {
		log.Fatal("Create", err)
	}
}
