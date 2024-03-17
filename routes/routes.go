package routes

import (
	"echo-api/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello this is echo")
	})

	e.GET("/employee", controllers.FetchAllEmployee)
	e.POST("/employee", controllers.StoreEmployee)
	e.PUT("/employee", controllers.UpdateEmployee)
	e.DELETE("/employee", controllers.DeleteEmployee)


	return e
}
