package controllers

import (
	"server/database"
	"server/entities"
	"server/repositories"
	"server/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GroupController interface {
	GetById(echo.Context) error
	GetAll(echo.Context) error
	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type groupController struct {
	groupService services.GroupService
}

func NewGroupController() GroupController {
	return &groupController{
		groupService: services.NewGroupService(
			repositories.NewGroupRepository(
				database.NewSqlHandler(
					database.GetDatabase(),
				),
			),
		),
	}
}

func (controller *groupController) GetById(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	group, err := controller.groupService.GetById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, group)
	return
}

func (controller *groupController) GetAll(c echo.Context) (err error) {
	groups, err := controller.groupService.GetAll()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, groups)
	return
}

func (controller *groupController) Create(c echo.Context) (err error) {
	u := entities.Group{}
	c.Bind(&u)
	group, err := controller.groupService.Create(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, group)
	return
}

func (controller *groupController) Update(c echo.Context) (err error) {
	u := entities.Group{}
	c.Bind(&u)
	group, err := controller.groupService.Update(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, group)
	return
}

func (controller *groupController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	group := entities.Group{
		ID: id,
	}
	err = controller.groupService.Delete(group)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, group)
	return
}
