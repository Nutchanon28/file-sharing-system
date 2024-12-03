package http

import (
	"github.com/Nutchanon28/file-sharing-system/internal/app/interfaces"
	"github.com/labstack/echo/v4"
)

func MapAppRoutes(
	appGroup *echo.Group,
	appHandler interfaces.AppHandlers,
) {
	appGroup.GET("/health", appHandler.Health())
}
