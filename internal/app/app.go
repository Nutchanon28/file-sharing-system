package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
}

func NewApp() *App {
	return &App{
		echo: echo.New(),
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

	// goroutine concurrent -> start the server without interrupting other stuff
	go func() {
		if err := a.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
			fmt.Println("shutting down the server")
		}
	}()

	<-ctx.Done()

	return nil
}
