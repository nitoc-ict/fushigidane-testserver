package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/coordinate", func(c echo.Context) error {
		address := c.QueryParam("address")
		if address == "" {
			c.JSON(http.StatusBadRequest, `{"status": "くえりください", "lat": 0, "long": 0}`)

			return nil
		}

		c.JSON(http.StatusOK, `{"status": "OK", "lat": 26.5263, "long": 128.031}`)

		return nil
	})

	return e
}
