package controllers

import (
	"server/database"
	"server/entities"
	"server/repositories"
	"server/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	Show(echo.Context) error
	Index(echo.Context) error
	Create(echo.Context) error
	Save(echo.Context) error
	Delete(echo.Context) error
}

type userController struct {
	userService services.UserService
}

func NewUserController(sqlHandler database.SqlHandler) UserController {
	return &userController{
		userService: services.NewUserService(
			repositories.NewUserRepository(sqlHandler),
		),
	}
}

func (controller *userController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.userService.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

func (controller *userController) Index(c echo.Context) (err error) {
	users, err := controller.userService.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

func (controller *userController) Create(c echo.Context) (err error) {
	u := entities.User{}
	c.Bind(&u)
	user, err := controller.userService.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *userController) Save(c echo.Context) (err error) {
	u := entities.User{}
	c.Bind(&u)
	user, err := controller.userService.Update(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *userController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := entities.User{
		ID: id,
	}
	err = controller.userService.DeleteById(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
