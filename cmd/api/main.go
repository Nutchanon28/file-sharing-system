package main

import (
	"fmt"

	"github.com/Nutchanon28/file-sharing-system/config"
	"github.com/Nutchanon28/file-sharing-system/internal/app"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v", err)
	}

	app := app.NewApp(config)
	if err := app.Run(); err != nil {
		fmt.Printf("server failed to run: %v", err)
	}
}
