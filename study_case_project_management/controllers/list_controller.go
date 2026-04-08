package controllers

import (
	"github.com/SatrioHalim/Go-x-React-Journey/models"
	"github.com/SatrioHalim/Go-x-React-Journey/services"
	"github.com/SatrioHalim/Go-x-React-Journey/utils"
	"github.com/gofiber/fiber/v3"
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