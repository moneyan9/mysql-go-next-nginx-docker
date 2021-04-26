package main

import (
	"server/infrastructure"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8000"))

	infrastructure.InitializeDatabase()
	infrastructure.InitializeRoots(e)
}
