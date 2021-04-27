package repositories

import (
	"server/infrastructure/entities"
	"server/infrastructure/interfaces"
)

type UserRepository struct {
	interfaces.ISqlHandler
}

func (repo *UserRepository) FindById(id int) (user entities.User, err error) {
	if err = repo.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users entities.Users, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) Store(u entities.User) (user entities.User, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(u entities.User) (user entities.User, err error) {
	if err = repo.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) DeleteById(user entities.User) (err error) {
	if err = repo.Delete(&user).Error; err != nil {
		return
	}
	return
}
