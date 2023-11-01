package main

import (
	"log"
	"refactoring/internal/app"
	"refactoring/internal/config"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app := app.NewApp(cfg)
	app.Run()
}