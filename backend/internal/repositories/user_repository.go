package repositories

import (
	"errors"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/pkg/db"
)

type IUserRepository interface {
	FindById(id uint) (models.User, error)
	FindByUsername(username string) (models.User, error)
	Create(userData data.UserData) (models.User, error)
	Update(userData data.UserData) (models.User, error)
}

type UserRepository struct {
	*db.IDbHandler
}

func (repository *UserRepository) FindById(id uint) (models.User, error) {
	var user models.User
	repository.DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)

	return user, nil
}

func (repository *UserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User

	result := repository.DB.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (repository *UserRepository) Create(userData data.UserData) (models.User, error) {
	var user models.User
	user.Name = userData.Name
	user.Username = userData.Username
	user.Password = userData.Password
	result := repository.DB.Where("username = ?", user.Username).FirstOrCreate(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.User{}, errors.New("user exists")
	}

	return user, nil
}

func (repository *UserRepository) Update(userData data.UserData) (models.User, error) {
	var user models.User

	user.Name = userData.Name
	user.Username = userData.Username
	user.Password = userData.Password

	if result := repository.DB.Where("id = ?", userData.Id).Updates(&user); result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
