package main

import (
	"staycation/app"
	"staycation/logger"
)

func main() {
	logger.Info("Starting our application...")
	app.Start()
}
