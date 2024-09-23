package controller

import (
	"github.com/gofiber/fiber/v2"
	"newserver/internal/usecase"
)

type CartController struct {
	useCase usecase.CartUseCase
}

func NewCartController(useCase usecase.CartUseCase) *CartController {
	return &CartController{useCase: useCase}
}

func (c *CartController) Get(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Get Cart")
}

func (c *CartController) Post(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Post Cart")
}

func (c *CartController) Put(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Put Cart")
}

func (c *CartController) Delete(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Delete Cart")
}
