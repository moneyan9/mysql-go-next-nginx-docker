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
	GetById(echo.Context) error
	GetAll(echo.Context) error
	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type userController struct {
	userService services.UserService
}

func NewUserController() UserController {
	return &userController{
		userService: services.NewUserService(
			repositories.NewUserRepository(
				database.NewSqlHandler(
					database.GetDatabase(),
				),
			),
		),
	}
}

func (controller *userController) GetById(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.userService.GetById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

func (controller *userController) GetAll(c echo.Context) (err error) {
	users, err := controller.userService.GetAll()
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
	user, err := controller.userService.Create(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *userController) Update(c echo.Context) (err error) {
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
	err = controller.userService.Delete(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
