package services

import (
	"server/infrastructure/entities"
	"server/services/interactors"
)

type UserService struct {
	UserInteractor interactors.UserInteractor
}

func (service *UserService) UserById(id int) (user entities.User, err error) {
	user, err = service.UserInteractor.FindById(id)
	return
}

func (service *UserService) Users() (users entities.Users, err error) {
	users, err = service.UserInteractor.FindAll()
	return
}

func (service *UserService) Add(u entities.User) (user entities.User, err error) {
	user, err = service.UserInteractor.Store(u)
	return
}

func (service *UserService) Update(u entities.User) (user entities.User, err error) {
	user, err = service.UserInteractor.Update(u)
	return
}

func (service *UserService) DeleteById(u entities.User) (err error) {
	err = service.UserInteractor.DeleteById(u)
	return
}
