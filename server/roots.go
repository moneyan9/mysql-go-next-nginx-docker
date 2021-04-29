package main

import (
	"server/controllers"

	"github.com/labstack/echo/v4"
)

func InitializeRoots(e *echo.Echo) {

	userController := controllers.NewUserController()
	e.GET("/api/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/api/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/api/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/api/users/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/api/users/:id", func(c echo.Context) error { return userController.Delete(c) })
}
