package controllers

import (
	"github.com/SatrioHalim/Go-x-React-Journey/models"
	"github.com/SatrioHalim/Go-x-React-Journey/services"
	"github.com/SatrioHalim/Go-x-React-Journey/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController{
	return &UserController{service: s}
}

func (c *UserController) Register(ctx fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.Bind().Body(user); err != nil {
		return utils.BadRequest(ctx,"Gagal Parsing Data",err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Registrasi Gagal",err.Error())
	}
	var userResp models.UserResponse
	_ = copier.Copy(&userResp,&user)
	return utils.Success(ctx,"Register Success",userResp)
}