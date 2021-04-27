package services

import (
	"server/infrastructure/entities"
	"server/infrastructure/repositories/interfaces"
)

type UserService struct {
	IUserRepository interfaces.IUserRepository
}

func (service *UserService) UserById(id int) (user entities.User, err error) {
	user, err = service.IUserRepository.FindById(id)
	return
}

func (service *UserService) Users() (users entities.Users, err error) {
	users, err = service.IUserRepository.FindAll()
	return
}

func (service *UserService) Add(u entities.User) (user entities.User, err error) {
	user, err = service.IUserRepository.Store(u)
	return
}

func (service *UserService) Update(u entities.User) (user entities.User, err error) {
	user, err = service.IUserRepository.Update(u)
	return
}

func (service *UserService) DeleteById(u entities.User) (err error) {
	err = service.IUserRepository.DeleteById(u)
	return
}
