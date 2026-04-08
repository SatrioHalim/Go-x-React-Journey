package controllers

import (
	"github.com/SatrioHalim/Go-x-React-Journey/models"
	"github.com/SatrioHalim/Go-x-React-Journey/services"
	"github.com/SatrioHalim/Go-x-React-Journey/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ListController struct {
	service services.ListService
}

func NewListController(s services.ListService) *ListController{
	return &ListController{service: s}
}

func (c *ListController) CreateList(ctx fiber.Ctx) error{
	list := new(models.List)
	if err := ctx.Bind().Body(list); err != nil {
		return utils.BadRequest(ctx, "Failed to read request",err.Error())
	}
	if err := c.service.Create(list); err != nil {
		return utils.BadRequest(ctx, "Failed to create list",err.Error())
	}
	return utils.Success(ctx,"List created successfully",list)
}

func (c *ListController) UpdateList(ctx fiber.Ctx) error{
	publicID := ctx.Params("id")
	list := new(models.List)
	if err := ctx.Bind().Body(list); err != nil{
		return utils.BadRequest(ctx, "Failed to parse data",err.Error())
	}
	
	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid ID",err.Error())
	}

	existingList, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List not found", err.Error())
	}
	list.InternalID = existingList.InternalID
	list.PublicID = existingList.PublicID

	if err := c.service.Update(list); err != nil {
		return utils.BadRequest(ctx, "Failed to update list", err.Error())
	}

	updatedList, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List not found", err.Error())
	}
	return utils.Success(ctx, "List updated successfully", updatedList)
}