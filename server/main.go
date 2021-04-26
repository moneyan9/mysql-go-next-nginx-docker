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

	infrastructure.InitializeDatabase()
	infrastructure.InitializeRoots(e)

	e.Logger.Fatal(e.Start(":8000"))

}
