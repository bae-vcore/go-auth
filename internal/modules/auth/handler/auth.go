package handler

import (
	"go-auth/internal/modules/auth/entity"
	"go-auth/internal/modules/auth/services"
	"go-auth/internal/modules/helper"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authHandler struct {
	service services.AuthServices
}

func NewAuthHandler(service services.AuthServices) *authHandler {
	return &authHandler{service: service}
}

func (s *authHandler) Login(c *fiber.Ctx) error {
	payload := new(entity.LoginReq)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something wrong",
			"data":    nil,
		})

	}

	user, err := s.service.Login(*payload)

	isMatch := helper.CheckPasswordHash(payload.Password, user.Password)

	if err != nil || !isMatch {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "email or password is wrong",
		})
	}

	token, err := helper.CreateToken(user.Email)

	if err != nil {
		log.Println("error while create token", err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "successfully login",
		"data": fiber.Map{
			"user":  user,
			"token": token,
		},
	})

}

func (s *authHandler) Register(c *fiber.Ctx) error {
	payload := new(entity.RegisterReq)

	if err := c.BodyParser(&payload); err != nil {
		// validasi
		return err
	}

	s.service.Register(*payload)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "successfully register",
	})

}
