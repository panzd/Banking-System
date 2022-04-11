package main

import (
	"github.com/Banking-System/app"
	"github.com/Banking-System/logger"
)

func main() {

	logger.Info("Starting our application...")
	// log.Println("Starting our application....")
	app.Start()
}
