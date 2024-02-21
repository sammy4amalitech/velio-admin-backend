package main

import (
	"github.com/joho/godotenv"
	"velio-admin-backend/config"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {

	// Now you can use userRepo to interact with the User table in your database
}
