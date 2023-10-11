package server

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
	"microapp-fiber-kit/config"
	"microapp-fiber-kit/internal/board"
	"microapp-fiber-kit/internal/user"
	router "microapp-fiber-kit/server/router"
)

func Api(
	conf *config.Config,
	lc fx.Lifecycle,
	api *FiberApiServer,
	boardSrv *board.BoardService,
	userSrv *user.UserService,
) {

	router.Route(api.Server, boardSrv, userSrv)

	router.SwaggerRoute(api.Server)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := api.Listen(conf.Host + ":" + conf.Port)
				if err != nil {
					log.Error(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return api.Shutdown()
		},
	})
}
