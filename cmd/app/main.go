package main

import (
	"log"

	"github.com/22Fariz22/trueconf/internal/app"
	"github.com/22Fariz22/trueconf/internal/config"
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
