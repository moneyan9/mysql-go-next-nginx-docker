package repositories

import (
	"server/database"
	"server/entities"
)

type UserRepository interface {
	GetById(id int) (entities.User, error)
	GetAll() (entities.Users, error)
	Create(entities.User) (entities.User, error)
	Update(entities.User) (entities.User, error)
	Delete(entities.User) error
}

type userRepository struct {
	sqlHandler database.SqlHandler
}

func NewUserRepository(sqlHandler database.SqlHandler) UserRepository {
	return &userRepository{sqlHandler: sqlHandler}
}

func (repo *userRepository) GetById(id int) (user entities.User, err error) {
	if err = repo.sqlHandler.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *userRepository) GetAll() (users entities.Users, err error) {
	if err = repo.sqlHandler.Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *userRepository) Create(u entities.User) (user entities.User, err error) {
	if err = repo.sqlHandler.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *userRepository) Update(u entities.User) (user entities.User, err error) {
	if err = repo.sqlHandler.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *userRepository) Delete(user entities.User) (err error) {
	if err = repo.sqlHandler.Delete(&user).Error; err != nil {
		return
	}
	return
}
