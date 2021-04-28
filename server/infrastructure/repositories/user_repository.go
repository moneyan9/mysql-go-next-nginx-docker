package repositories

import (
	"server/infrastructure"
	"server/infrastructure/entities"
)

type UserRepository interface {
	FindById(id int) (entities.User, error)
	FindAll() (entities.Users, error)
	Store(entities.User) (entities.User, error)
	Update(entities.User) (entities.User, error)
	DeleteById(entities.User) error
}

type userRepository struct {
	sqlHandler infrastructure.SqlHandler
}

func NewUserRepository(sqlHandler infrastructure.SqlHandler) UserRepository {
	return &userRepository{sqlHandler: sqlHandler}
}

func (repo *userRepository) FindById(id int) (user entities.User, err error) {
	if err = repo.sqlHandler.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *userRepository) FindAll() (users entities.Users, err error) {
	if err = repo.sqlHandler.Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *userRepository) Store(u entities.User) (user entities.User, err error) {
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

func (repo *userRepository) DeleteById(user entities.User) (err error) {
	if err = repo.sqlHandler.Delete(&user).Error; err != nil {
		return
	}
	return
}
