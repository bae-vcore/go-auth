package handler

import (
	"go-auth/internal/modules/user/entity"
	"go-auth/internal/modules/user/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetAllUsers(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()

	if err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "message": "failed to get all users", "data": nil})
	}

	return c.JSON(fiber.Map{"success": true, "messages": "successfully get all users", "data": users})
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {

	newUser := new(entity.User)

	if err := c.BodyParser(&newUser); err != nil {
		// validasi
		return err
	}

	err := h.service.CreateUser(newUser)

	if err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "message": "failed to create user", "data": nil})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "messages": "successfully create user"})
}

func (h *userHandler) GetUserByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := h.service.GetUserByID(id)

	if err != nil {
		log.Println("error while get user by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "failed to get user by id",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "successfully get user",
		"data":    user,
	})
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := h.service.DeleteUser(id)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "successfully delete user",
	})
}
