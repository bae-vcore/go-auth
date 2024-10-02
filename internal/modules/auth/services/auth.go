package services

import (
	"errors"
	"go-auth/internal/modules/auth/entity"
	"go-auth/internal/modules/helper"
	userEntity "go-auth/internal/modules/user/entity"
	"go-auth/internal/modules/user/repository"
)

type AuthServices interface {
	Login(payload entity.LoginReq) (userEntity.User, error)
	Register(payload entity.RegisterReq) error
}

type authServices struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) *authServices {
	return &authServices{repository: repository}
}

func (s *authServices) Login(payload entity.LoginReq) (userEntity.User, error) {
	// check if email exists
	user, err := s.repository.GetUserByEmail(payload.Email)

	if err != nil {
		return user, errors.New("email or password is wrong")
	}

	// check if password is match
	if helper.CheckPasswordHash(payload.Password, user.Password) {
		return user, err
	}
	return user, nil
}

func (s *authServices) Register(payload entity.RegisterReq) error {

	hashedPassword, _ := helper.HashPassword(payload.Password)

	newUser := &userEntity.User{
		Email:    payload.Email,
		Password: hashedPassword,
		Name:     payload.Name,
	}

	err := s.repository.CreateUser(newUser)

	return err

}
