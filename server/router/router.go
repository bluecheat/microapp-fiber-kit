package router

import (
	"github.com/gofiber/fiber/v2"
)

const (
	V1 = "/v1"
)

func NoAuthRoute(
	api *fiber.App,
) {

	v1 := api.Group(V1)

	v1.Get("/board", func(ctx *fiber.Ctx) error {

		return nil
	})
}
