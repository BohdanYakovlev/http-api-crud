package controller

import (
	"github.com/gofiber/fiber/v2"
	"httpApiTest/api"
)

// checkController handles the logic for checking the application's status.
type userController struct {
	// Add any necessary initialization for use cases or other components here.
}

// ICheckController defines the interface for the checkController.
type IUserController interface {
	Execute(c *fiber.Ctx) error
}

// NewCheckController creates a new instance of checkController.
func NewUserController() IUserController {
	return &userController{}
}

// Execute implements the ICheckController interface.
func (*userController) Execute(c *fiber.Ctx) error {
	var apiRes api.User
	return c.JSON(apiRes)
}
