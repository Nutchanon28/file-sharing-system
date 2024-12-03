package app

import (
	appHttp "github.com/Nutchanon28/file-sharing-system/internal/app/delivery/http"
	"github.com/labstack/echo/v4/middleware"
)

func (s *App) MapHandlers() error {
	s.echo.Use(middleware.CORS())

	baseGroup := s.echo.Group("")

	appHandlers := appHttp.NewAppHandlers()

	appHttp.MapAppRoutes(baseGroup, appHandlers)

	return nil
}
