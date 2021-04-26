package interactors

import "server/infrastructure/entities"

type UserInteractor interface {
	FindById(id int) (entities.User, error)
	FindAll() (entities.Users, error)
	Store(entities.User) (entities.User, error)
	Update(entities.User) (entities.User, error)
	DeleteById(entities.User) error
}
