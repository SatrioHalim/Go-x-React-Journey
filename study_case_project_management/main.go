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

	// User handling
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Board handling
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo,userRepo,boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.Setup(app, userController,boardController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port :",port)
	log.Fatal(app.Listen(":"+port))

}