package controllers

import (
	"backend-api/exceptions"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		exceptions.Handle("Error loading .env file")
	}
}
