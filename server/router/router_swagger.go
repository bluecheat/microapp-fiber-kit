package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	// _ "microapp-fiber-kit/docs"
)

func SwaggerRoute(
	api *fiber.App,
) {
	api.Get("/swagger/*", swagger.HandlerDefault) // default
}
