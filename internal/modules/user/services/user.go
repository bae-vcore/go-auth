package services

import (
	"go-auth/internal/modules/user/entity"
	"go-auth/internal/modules/user/repository"
)

type UserService interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id int) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user *entity.User) error
	DeleteUser(id int) error
	UpdateUser(user entity.User) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(user *entity.User) error {
	err := s.repository.CreateUser(user)

	return err
}

func (s *userService) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)

	return err
}

func (s *userService) GetUserByID(id int) (entity.User, error) {
	user, err := s.repository.GetUserByID(id)

	return user, err
}

func (s *userService) GetUserByEmail(email string) (entity.User, error) {
	user, err := s.repository.GetUserByEmail(email)

	return user, err
}

func (s *userService) UpdateUser(user entity.User) error {
	err := s.repository.UpdateUser(user)

	return err
}

func (s *userService) GetAllUsers() ([]entity.User, error) {
	users, err := s.repository.GetAllUsers()

	return users, err
}
