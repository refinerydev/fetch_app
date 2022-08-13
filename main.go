package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "success"})
	})

	e.Logger.Fatal(e.Start(":8081"))
}
