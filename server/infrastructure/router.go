package infrastructure

import (
	"server/interfaces/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userController := controllers.NewUserController(NewSqlHandler())

	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	e.Logger.Fatal(e.Start(":8000"))
}
