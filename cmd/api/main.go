package main

import (
	"fmt"

	"github.com/Nutchanon28/file-sharing-system/internal/app"
)

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		fmt.Printf("server failed to run: %v", err)
	}
}
