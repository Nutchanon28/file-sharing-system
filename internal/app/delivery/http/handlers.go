package http

import (
	"net/http"

	"github.com/Nutchanon28/file-sharing-system/internal/app/interfaces"
	"github.com/labstack/echo/v4"
)

type AppHandlers struct {
}

func NewAppHandlers() interfaces.AppHandlers {
	return &AppHandlers{}
}

func (a *AppHandlers) Health() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "OK",
		})
	}
}
