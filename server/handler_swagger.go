package server

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "microapp-fiber-kit/docs"
)

func SwaggerRoute(
	api *FiberApiServer,
) {
	api.server.Get("/swagger/*", swagger.HandlerDefault) // default
}
