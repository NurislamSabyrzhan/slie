package controller

import (
	"github.com/gofiber/fiber/v2"
	"newserver/internal/usecase"
)

type usersController struct {
	useCase usecase.usersUseCase
}

func NewusersController(useCase usecase.usersUseCase) *usersController {
	return &usersController{useCase: useCase}
}

func (c *usersController) Get(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Get users")
}

func (c *usersController) Post(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Post users")
}

func (c *usersController) Put(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Put users")
}

func (c *usersController) Delete(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Delete users")
}
