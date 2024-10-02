package routes

import (
	"go-auth/internal/middleware"
	ah "go-auth/internal/modules/auth/handler"
	uh "go-auth/internal/modules/user/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	UserHandler uh.UserHandler
	AuthHandler ah.AuthHandler
}

func (h *Handler) ServeAndListen() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// root handler
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "learn create an auth with go",
		})
	})

	// user
	h.NewUserRoutes(v1)
	h.NewAuthRoutes(v1)

	app.Listen(":3001")

}

func (h *Handler) NewUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Get("/all", middleware.Protected, h.UserHandler.GetAllUsers)
	user.Post("/create", h.UserHandler.CreateUser)
	user.Get("/:id", h.UserHandler.GetUserByID)
	user.Delete("/:id", h.UserHandler.DeleteUser)
}

func (h *Handler) NewAuthRoutes(router fiber.Router) {
	user := router.Group("/auth")
	user.Post("/register", h.AuthHandler.Register)
	user.Post("/login", h.AuthHandler.Login)
}
