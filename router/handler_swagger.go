package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "microapp-fiber-kit/docs"
	"microapp-fiber-kit/server"
)

func SwaggerRoute(
	api *server.FiberApiServer,
) {
	api.Server.Get("/swagger/*", swagger.HandlerDefault) // default
}
