package repository

import (
	"go-auth/internal/modules/user/entity"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id int) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user *entity.User) error
	DeleteUser(id int) error
	UpdateUser(user entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	result := r.db.Create(&user)

	if result.Error != nil {
		log.Println("error while create user", result.Error)
	}

	return result.Error
}

func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	var user []entity.User

	result := r.db.Find(&user)
	if result.Error != nil {
		log.Println("error while get all users", result.Error)
	}

	return user, result.Error
}

func (r *userRepository) GetUserByID(id int) (entity.User, error) {
	var user entity.User

	result := r.db.First(&user, id)

	return user, result.Error
}

func (r *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User

	result := r.db.Where("email = ?", email).First(&user)

	return user, result.Error
}

func (r *userRepository) DeleteUser(id int) error {
	var user entity.User

	result := r.db.Where("ID = ?", id).Delete(&user)

	return result.Error
}

func (r *userRepository) UpdateUser(user entity.User) error {
	result := r.db.Where("ID = ?", user.ID).Save(user)

	return result.Error
}
