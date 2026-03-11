package main

import (
	"log"

	"github.com/SatrioHalim/Go-x-React-Journey/config"
	"github.com/SatrioHalim/Go-x-React-Journey/controllers"
	"github.com/SatrioHalim/Go-x-React-Journey/database/seed"
	"github.com/SatrioHalim/Go-x-React-Journey/repositories"
	"github.com/SatrioHalim/Go-x-React-Journey/routes"
	"github.com/SatrioHalim/Go-x-React-Journey/services"
	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port :",port)
	log.Fatal(app.Listen(":"+port))

}