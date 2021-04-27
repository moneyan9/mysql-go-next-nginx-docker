package main

import (
	"server/infrastructure"
	"server/roots"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	infrastructure.InitializeDatabase()
	roots.InitializeRoots(e)

	e.Logger.Fatal(e.Start(":8000"))

}
