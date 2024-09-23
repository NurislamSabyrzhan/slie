package router

import (
	"github.com/gofiber/fiber/v2"
	"newserver/internal/controller"
	"newserver/internal/usecase"
)

func NewCartRouter(app *fiber.App, useCase usecase.CartUseCase) {
	ctrl := controller.NewCartController(useCase)
	cart := app.Group("/cart")

	cart.Get("/", ctrl.Get)
	cart.Post("/", ctrl.Post)
	cart.Put("/:id", ctrl.Put)
	cart.Delete("/:id", ctrl.Delete)
}
