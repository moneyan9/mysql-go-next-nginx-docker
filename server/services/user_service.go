package services

import (
	"server/entities"
	"server/repositories"
)

type UserService interface {
	GetById(id int) (entities.User, error)
	GetAll() (entities.Users, error)
	Create(entities.User) (entities.User, error)
	Update(entities.User) (entities.User, error)
	Delete(entities.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (service *userService) GetById(id int) (user entities.User, err error) {
	user, err = service.userRepository.GetById(id)
	return
}

func (service *userService) GetAll() (users entities.Users, err error) {
	users, err = service.userRepository.GetAll()
	return
}

func (service *userService) Create(u entities.User) (user entities.User, err error) {
	user, err = service.userRepository.Create(u)
	return
}

func (service *userService) Update(u entities.User) (user entities.User, err error) {
	user, err = service.userRepository.Update(u)
	return
}

func (service *userService) Delete(u entities.User) (err error) {
	err = service.userRepository.Delete(u)
	return
}
