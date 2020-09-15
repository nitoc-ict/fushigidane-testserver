package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kr/pretty"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type routeRequest struct {
	Lat   float64 `json:"latitude"`
	Long  float64 `json:"longitude"`
	Label string  `json:"label"`
}

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

	e.GET("/route", func(c echo.Context) error {
		var routeData routeRequest
		err := c.Bind(&routeData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, `{"status": "error bind json"}`)
			return nil
		}

		jsonFile, err := ioutil.ReadFile("./data.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, `{"status": "error read file"}`)
			return nil
		}
		route := []routeRequest{}

		err = json.Unmarshal(jsonFile, &route)

		pretty.Println(route)

		c.JSON(http.StatusOK, route)

		return nil
	})

	return e
}
