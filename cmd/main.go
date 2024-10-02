package main

import (
	"go-auth/config"
	ue "go-auth/internal/modules/user/entity"

	ah "go-auth/internal/modules/auth/handler"
	uh "go-auth/internal/modules/user/handler"

	ur "go-auth/internal/modules/user/repository"

	as "go-auth/internal/modules/auth/services"
	us "go-auth/internal/modules/user/services"

	"go-auth/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env file")
	}

	// init db
	db := config.InitDB()
	db.AutoMigrate(&ue.User{})

	// init repo
	userRepo := ur.NewUserRepository(db)

	// init service
	userSvc := us.NewUserService(userRepo)
	authSvc := as.NewAuthService(userRepo)

	// init handler
	userHandler := uh.NewUserHandler(userSvc)
	authHandler := ah.NewAuthHandler(authSvc)

	// server and listen
	handlers := routes.Handler{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}

	handlers.ServeAndListen()

}
