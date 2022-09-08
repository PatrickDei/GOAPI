package main

import (
	"os"
	"staycation/app"
	"staycation/logger"
)

func main() {
	logger.Info("Listening on port " + os.Getenv("SERVER_PORT"))
	app.Start()
}
