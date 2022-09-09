package main

import (
	"os"
	"staycation/app"
	"staycation/logger"
)

//SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWORD=root DB_ADDRESS=localhost DB_PORT=3306 DB_NAME=Staycation go run main.go
func main() {
	logger.Info("Listening on port " + os.Getenv("SERVER_PORT"))
	app.Start()
}
