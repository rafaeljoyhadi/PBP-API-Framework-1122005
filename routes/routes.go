package routes

import "net/http"
import "github.com/labstack/echo/v4"

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello this is echo")
	})

	return e
}