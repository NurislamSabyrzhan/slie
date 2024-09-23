package router

import (
	"github.com/gofiber/fiber/v2"
	"newserver/internal/controller"
	"newserver/internal/usecase"
)

func NewusersRouter(app *fiber.App, useCase usecase.usersUseCase) {
	ctrl := controller.NewusersController(useCase)
	users := app.Group("/users")

	users.Get("/", ctrl.Get)
	users.Post("/", ctrl.Post)
	users.Put("/:id", ctrl.Put)
	users.Delete("/:id", ctrl.Delete)
}
