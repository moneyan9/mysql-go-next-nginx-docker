package main

import (
	"server/controllers"

	"github.com/labstack/echo/v4"
)

func InitializeRoots(e *echo.Echo) {

	groupController := controllers.NewGroupController()
	e.GET("/api/groups", func(c echo.Context) error { return groupController.GetAll(c) })
	e.GET("/api/groups/:id", func(c echo.Context) error { return groupController.GetById(c) })
	e.POST("/api/groups", func(c echo.Context) error { return groupController.Create(c) })
	e.PUT("/api/groups/:id", func(c echo.Context) error { return groupController.Update(c) })
	e.DELETE("/api/groups/:id", func(c echo.Context) error { return groupController.Delete(c) })

	userController := controllers.NewUserController()
	e.GET("/api/users", func(c echo.Context) error { return userController.GetAll(c) })
	e.GET("/api/users/:id", func(c echo.Context) error { return userController.GetById(c) })
	e.POST("/api/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/api/users/:id", func(c echo.Context) error { return userController.Update(c) })
	e.DELETE("/api/users/:id", func(c echo.Context) error { return userController.Delete(c) })
}
