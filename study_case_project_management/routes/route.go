package routes

import (
	"log"

	"github.com/SatrioHalim/Go-x-React-Journey/controllers"
	"github.com/SatrioHalim/Go-x-React-Journey/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App, 
	uc *controllers.UserController,
	bc *controllers.BoardController,
	){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Post("/v1/auth/register",uc.Register)
	app.Post("/v1/auth/login",uc.Login)
	
	// JWT Protected routes : biar ga sembarangan orang bisa akses
	api := app.Group("/api/v1", middleware.JWTAuth())

	userGroup := api.Group("/users")
	userGroup.Get("/page",uc.GetUserPagination) // /api/v1/users/:id
	userGroup.Get("/:id",uc.GetUser) // /api/v1/users/:id
	userGroup.Put("/:id",uc.UpdateUser)
	userGroup.Delete("/:id",uc.DeleteUser)

	boardGroup := api.Group("/boards")
	boardGroup.Post("/",bc.CreateBoard)
}