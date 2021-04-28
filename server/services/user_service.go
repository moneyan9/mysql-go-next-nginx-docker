package services

import (
	"server/infrastructure/entities"
	"server/infrastructure/repositories"
)

type UserService interface {
	UserById(id int) (entities.User, error)
	Users() (entities.Users, error)
	Add(entities.User) (entities.User, error)
	Update(entities.User) (entities.User, error)
	DeleteById(entities.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (service *userService) UserById(id int) (user entities.User, err error) {
	user, err = service.userRepository.FindById(id)
	return
}

func (service *userService) Users() (users entities.Users, err error) {
	users, err = service.userRepository.FindAll()
	return
}

func (service *userService) Add(u entities.User) (user entities.User, err error) {
	user, err = service.userRepository.Store(u)
	return
}

func (service *userService) Update(u entities.User) (user entities.User, err error) {
	user, err = service.userRepository.Update(u)
	return
}

func (service *userService) DeleteById(u entities.User) (err error) {
	err = service.userRepository.DeleteById(u)
	return
}
