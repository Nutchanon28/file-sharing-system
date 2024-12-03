package interfaces

import "github.com/labstack/echo/v4"

type AppHandlers interface {
	Health() echo.HandlerFunc
}
