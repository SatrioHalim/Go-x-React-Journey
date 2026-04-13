package controllers

import (
	"time"

	"github.com/SatrioHalim/Go-x-React-Journey/models"
	"github.com/SatrioHalim/Go-x-React-Journey/services"
	"github.com/SatrioHalim/Go-x-React-Journey/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type CardController struct {
	service services.CardService
}

func NewCardController(service services.CardService) *CardController {
	return &CardController{service}
}

func (c *CardController) CreateCard(ctx fiber.Ctx) error{
	type CreateCardRequest struct {
		ListPublicID string `json:"list_id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		DueDate time.Time `json:"due_date"`
		Position int `json:"position"`
	}

	var req CreateCardRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return utils.BadRequest(ctx, "Failed to fetch data",err.Error())
	}
	card := &models.Card{
		Title: req.Title,
		Description: req.Description,
		DueDate: &req.DueDate,
		Position: req.Position,
	}
	if err := c.service.Create(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx, "Failed to create card",err.Error())
	}
	return utils.Success(ctx, "Card created successfully",card)
}

func (c *CardController) UpdateCard(ctx fiber.Ctx) error {
	publicID := ctx.Params("id")

	type updateCardRequest struct{
		ListPublicID string `json:"list_id"`
		Title string `json:"title"`
		Description string `json:"description"`
		DueDate *time.Time `json:"due_date"`
		Position int `json:"position"`
	}

	var req updateCardRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return utils.BadRequest(ctx, "Failed to parsing data",err.Error())
	}

	if _,err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid id",err.Error())
	}

	card := &models.Card{
		Title: req.Title,
		Description: req.Description,
		DueDate: req.DueDate,
		Position: req.Position,
		PublicID: uuid.MustParse(publicID),
	}

	if err := c.service.Update(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx,"Failed to update data",err.Error())
	}

	return utils.Success(ctx,"Card updated successfully",card) 
}