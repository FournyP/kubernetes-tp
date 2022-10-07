package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}
}
