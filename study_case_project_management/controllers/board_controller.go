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

func (c *BoardController) UpdateBoard(ctx fiber.Ctx) error{
	publicID := ctx.Params("id")
	board := new(models.Board)
	if err := ctx.Bind().Body(board); err != nil {
		return utils.BadRequest(ctx,"Failed to parsing data",err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx,"Invalid ID",err.Error())
	}
	existingBoard, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx,"Board not found",err.Error())
	}
	board.InternalID = existingBoard.InternalID
	board.PublicID = existingBoard.PublicID

	if err := c.service.Update(board); err != nil{
		return utils.BadRequest(ctx,"Failed to update board",err.Error())
	}
	return utils.Success(ctx,"Board updated successfully",board)
}

func (c *BoardController) AddBoardMembers(ctx fiber.Ctx) error{
	publicID := ctx.Params("id")
	var userIDs []string

	if err := ctx.Bind().Body(&userIDs); err != nil {
		return utils.BadRequest(ctx,"Failed to parsing data",err.Error())
	}

	if err := c.service.AddMembers(publicID, userIDs); err != nil {
		return utils.BadRequest(ctx,"Failed to add members",err.Error())
	}
	
	return utils.Success(ctx,"Members added successfully",nil)
}