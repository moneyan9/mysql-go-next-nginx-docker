package repositories

import (
	"server/database"
	"server/entities"
)

type GroupRepository interface {
	GetById(id int) (entities.Group, error)
	GetAll() (entities.Groups, error)
	Create(entities.Group) (entities.Group, error)
	Update(entities.Group) (entities.Group, error)
	Delete(entities.Group) error
}

type groupRepository struct {
	sqlHandler database.SqlHandler
}

func NewGroupRepository(sqlHandler database.SqlHandler) GroupRepository {
	return &groupRepository{sqlHandler: sqlHandler}
}

func (repo *groupRepository) GetById(id int) (group entities.Group, err error) {
	if err = repo.sqlHandler.Find(&group, id).Error; err != nil {
		return
	}
	return
}

func (repo *groupRepository) GetAll() (groups entities.Groups, err error) {
	if err = repo.sqlHandler.Find(&groups).Error; err != nil {
		return
	}
	return
}

func (repo *groupRepository) Create(u entities.Group) (group entities.Group, err error) {
	if err = repo.sqlHandler.Create(&u).Error; err != nil {
		return
	}
	group = u
	return
}

func (repo *groupRepository) Update(u entities.Group) (group entities.Group, err error) {
	if err = repo.sqlHandler.Save(&u).Error; err != nil {
		return
	}
	group = u
	return
}

func (repo *groupRepository) Delete(group entities.Group) (err error) {
	if err = repo.sqlHandler.Delete(&group).Error; err != nil {
		return
	}
	return
}
