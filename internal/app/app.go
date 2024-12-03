package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/Nutchanon28/file-sharing-system/config"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo   *echo.Echo
	config *config.Config
}

// TODO: can't you just import the config? It's not like there's any other config anyways?
func NewApp(config *config.Config) *App {
	return &App{
		echo:   echo.New(),
		config: config,
	}
}

func (a *App) Run() error {
	if err := a.MapHandlers(); err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// prevent the server from stopping before go echo starts
	// since go echo start with concurrency
	defer stop()

	serverUrl := fmt.Sprintf(":%d", a.config.Port)

	// goroutine concurrent -> start the server without interrupting other stuff
	go func() {
		if err := a.echo.Start(serverUrl); err != nil && err != http.ErrServerClosed {
			fmt.Println("shutting down the server")
		}
	}()

	<-ctx.Done()

	return nil
}
