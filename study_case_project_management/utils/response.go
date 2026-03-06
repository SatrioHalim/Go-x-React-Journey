package utils

import "github.com/gofiber/fiber/v3"

/*
	Template Standard response
	{
		"status" : "success",
		"response_code" : 200,
		"message" : "Login successful",
		"data" : {}
	}
*/

type Response struct {
	Status string `json:"status"`
	ResponseCode int `json:"response_code"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// dipake sbagai informasi sukses yg masuk ke server
func success(c fiber.Ctx, message string, data interface{}) error{
	return c.Status(fiber.StatusOK).JSON(Response{
		Status: "Success",
		ResponseCode: fiber.StatusOK,
		Message: message,
		Data: data,
	})
}

func created(c fiber.Ctx, message string, data interface{}) error{
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status: "Created",
		ResponseCode: fiber.StatusCreated,
		Message: message,
		Data: data,
	})
}

func badRequest(c fiber.Ctx, message string, data interface{}, err string) error{
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status: "Error Bad Request",
		ResponseCode: fiber.StatusBadRequest,
		Message: message,
		Data: data,
		Error: err,
	})
}

func notFound(c fiber.Ctx, message string, data interface{},err string) error{
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status: "Error Not Found",
		ResponseCode: fiber.StatusNotFound,
		Message: message,
		Data: data,
		Error: err,
	})
}