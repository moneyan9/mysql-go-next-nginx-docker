package controllers

import (
	"server/helpers"
	"server/infrastructure/entities"
	"server/infrastructure/interactors"
	"server/infrastructure/repositories"
	"server/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service services.UserService
}

func NewUserController(sqlHandler interactors.SqlInteractor) *UserController {
	return &UserController{
		Service: services.UserService{
			UserInteractor: &repositories.UserRepository{
				SqlInteractor: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Service.UserById(id)
	if err != nil {
		c.JSON(500, helpers.NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Service.Users()
	if err != nil {
		c.JSON(500, helpers.NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	u := entities.User{}
	c.Bind(&u)
	user, err := controller.Service.Add(u)
	if err != nil {
		c.JSON(500, helpers.NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Save(c echo.Context) (err error) {
	u := entities.User{}
	c.Bind(&u)
	user, err := controller.Service.Update(u)
	if err != nil {
		c.JSON(500, helpers.NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := entities.User{
		ID: id,
	}
	err = controller.Service.DeleteById(user)
	if err != nil {
		c.JSON(500, helpers.NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
