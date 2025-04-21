package services

import (
	"errors"
	"fmt"
	"time"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/dto/requests"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(request *requests.Login) (data.JwtToken, error)
	Profile(userId uint) (data.UserProfileData, error)
	Register(request *requests.Register) (data.JwtToken, error)
	GenerateToken(userId string) (string, error)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

type AuthService struct {
	UserService IUserService
}

func (service *AuthService) GenerateToken(userId string) (string, error) {
	secretKey := viper.Get("APP_SECRET_KEY").(string)
	secretKeyByte := []byte(secretKey)
	iss := fmt.Sprintf("%v:%v", "jwtservice", "3321")
	claims := jwt.MapClaims{
		"sub":   userId,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
		"iss":   iss,
		"roles": []string{"admin"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = "sim2"
	signedToken, _ := token.SignedString(secretKeyByte)

	return signedToken, nil

}

func (service *AuthService) Register(request *requests.Register) (data.JwtToken, error) {

	password, _ := HashPassword(request.Password)

	resUser, err := service.UserService.Create(data.UserData{
		Name:     request.Name,
		Username: request.Username,
		Password: password,
	})

	if err != nil {
		return data.JwtToken{}, err
	}

	resToken, _ := service.GenerateToken(fmt.Sprintf("%d", resUser.Id))

	return data.JwtToken{
		User: data.UserProfileData{
			Id:       resUser.Id,
			Name:     resUser.Name,
			Username: resUser.Username,
		},
		Token: resToken,
	}, nil
}

func (service *AuthService) Login(request *requests.Login) (data.JwtToken, error) {

	resUser, err := service.UserService.FindByUsername(request.Username)
	if err != nil {
		return data.JwtToken{}, errors.New("credential wrong")
	}

	if !checkPassword(resUser.Password, request.Password) {
		return data.JwtToken{}, errors.New("credential wrong")
	}

	resToken, _ := service.GenerateToken(fmt.Sprintf("%d", resUser.Id))
	return data.JwtToken{
		User: data.UserProfileData{
			Id:       resUser.Id,
			Name:     resUser.Name,
			Username: resUser.Username,
		},
		Token: resToken,
	}, nil
}

func (service *AuthService) Profile(userId uint) (data.UserProfileData, error) {

	resUser, err := service.UserService.FindById(userId)
	if err != nil {
		return data.UserProfileData{}, err
	}

	return data.UserProfileData{
		Id:       resUser.Id,
		Name:     resUser.Name,
		Username: resUser.Username,
	}, nil
}
