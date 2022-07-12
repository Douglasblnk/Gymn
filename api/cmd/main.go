package main

import (
	"gymn/internal/database"
	server "gymn/v1"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load config:", err)
	}

	if _, err := database.InitDB(); err != nil {
		log.Fatal("cannot connect to DB", err)
	}
}

func main() {
	server.InitServer()
}
