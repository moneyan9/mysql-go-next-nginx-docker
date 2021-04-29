package services

import (
	"server/entities"
	"server/repositories"
)

type GroupService interface {
	GetById(id int) (entities.Group, error)
	GetAll() (entities.Groups, error)
	Create(entities.Group) (entities.Group, error)
	Update(entities.Group) (entities.Group, error)
	Delete(entities.Group) error
}

type groupService struct {
	groupRepository repositories.GroupRepository
}

func NewGroupService(groupRepository repositories.GroupRepository) GroupService {
	return &groupService{groupRepository: groupRepository}
}

func (service *groupService) GetById(id int) (group entities.Group, err error) {
	group, err = service.groupRepository.GetById(id)
	return
}

func (service *groupService) GetAll() (groups entities.Groups, err error) {
	groups, err = service.groupRepository.GetAll()
	return
}

func (service *groupService) Create(u entities.Group) (group entities.Group, err error) {
	group, err = service.groupRepository.Create(u)
	return
}

func (service *groupService) Update(u entities.Group) (group entities.Group, err error) {
	group, err = service.groupRepository.Update(u)
	return
}

func (service *groupService) Delete(u entities.Group) (err error) {
	err = service.groupRepository.Delete(u)
	return
}
