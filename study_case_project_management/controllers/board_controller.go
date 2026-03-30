package controllers

import (
	"fmt"

	"github.com/SatrioHalim/Go-x-React-Journey/models"
	"github.com/SatrioHalim/Go-x-React-Journey/services"
	"github.com/SatrioHalim/Go-x-React-Journey/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController{
	return &BoardController{service: s}
}

func (c *BoardController) CreateBoard(ctx fiber.Ctx) error{
	var userID uuid.UUID
	var err error
	board := new(models.Board)
	claims,ok := ctx.Locals("user").(jwt.MapClaims)
	
	// debug
	userVal := ctx.Locals("user")
	fmt.Printf("DEBUG: user type = %T, value = %v\n", userVal, userVal)
	
	if !ok {
		return utils.Unauthorized(ctx,"Unauthorized","Invalid or missing token claims")
	}

	if err := ctx.Bind().Body(board); err != nil {
		return utils.BadRequest(ctx,"Failed to read request",err.Error())
	}

	userID, err = uuid.Parse(claims["pub_id"].(string))
	if err != nil{
		return utils.BadRequest(ctx,"Failed to read request",err.Error())
	}
	board.OwnerPublicID = userID

	if err := c.service.Create(board); err != nil {
		return utils.BadRequest(ctx,"Failed to save data",err.Error())

	}
	return utils.Success(ctx,"Board created successfully",board)
}