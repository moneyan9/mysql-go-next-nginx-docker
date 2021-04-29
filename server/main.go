package main

import (
	"server/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.InitializeDatabase()
	InitializeRoots(e)

	e.Logger.Fatal(e.Start(":8000"))

}
