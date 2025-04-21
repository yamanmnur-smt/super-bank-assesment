package services

import (
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/repositories"
)

type IUserService interface {
	FindById(id uint) (data.UserData, error)
	FindByUsername(username string) (data.UserData, error)
	Create(userData data.UserData) (data.UserData, error)
	Update(userData data.UserData) (data.UserData, error)
}

// Denpedency Injection for Repository
type UserService struct {
	Repository repositories.IUserRepository
}

func (service *UserService) FindById(id uint) (data.UserData, error) {

	var result data.UserData

	user, err := service.Repository.FindById(id)
	if err != nil {
		return data.UserData{}, err
	}

	result.Id = user.ID
	result.Name = user.Name
	result.Username = user.Username
	result.Password = user.Password

	return result, nil
}

func (service *UserService) FindByUsername(username string) (data.UserData, error) {

	var result data.UserData

	user, err := service.Repository.FindByUsername(username)
	if err != nil {
		return data.UserData{}, err
	}

	result.Id = user.ID
	result.Name = user.Name
	result.Username = user.Username
	result.Password = user.Password

	return result, nil
}

func (service *UserService) Create(userData data.UserData) (data.UserData, error) {

	var result data.UserData

	user, err := service.Repository.Create(userData)
	if err != nil {
		return data.UserData{}, err
	}

	result.Id = user.ID
	result.Name = user.Name
	result.Username = user.Username
	result.Password = user.Password

	return result, nil
}

func (service *UserService) Update(userData data.UserData) (data.UserData, error) {

	var result data.UserData

	user, err := service.Repository.Update(userData)
	if err != nil {
		return data.UserData{}, err
	}

	result.Id = user.ID
	result.Name = user.Name
	result.Username = user.Username
	result.Password = user.Password

	return result, nil
}
