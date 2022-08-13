package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/refinerydev/fetch_app/app"
)

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	app.NewRouter(e)

	server := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	e.Logger.Fatal(e.Start(server))
}
