package main

import (
	"server/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	configs.InitializeDatabase()
	configs.InitializeRoots(e)

	e.Logger.Fatal(e.Start(":8000"))

}
