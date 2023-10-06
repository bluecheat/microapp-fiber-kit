package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/fx"
	"gopkg.in/natefinch/lumberjack.v2"
	"microapp-fiber-kit/config"
	"microapp-fiber-kit/internal/board"
	"strings"
)

type FiberApiServer struct {
	conf     *config.Config
	server   *fiber.App
	boardSrv *board.BoardService
}

func NewFiberApiServer(conf *config.Config, boardSrv *board.BoardService) *FiberApiServer {
	server := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorHandler,
	})
	server.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(conf.Cors.Origins, ","),
		AllowMethods:     strings.Join(conf.Cors.Methods, ","),
		AllowHeaders:     strings.Join(conf.Cors.Headers, ","),
		AllowCredentials: conf.Cors.Credentials,
	}))
	server.Use(requestid.New(requestid.Config{
		Next:       nil,
		Header:     fiber.HeaderXRequestID,
		Generator:  utils.UUID,
		ContextKey: "rid",
	}))

	if conf.Log.Type == "file" {
		server.Use(logger.New(logger.Config{
			Output: &lumberjack.Logger{
				Filename:   conf.Log.Filename,
				MaxSize:    20, // megabytes
				MaxBackups: 3,
				MaxAge:     28, // days
			},
		}))
	} else {
		server.Use(logger.New())
	}
	server.Use(recover.New())

	return &FiberApiServer{
		server:   server,
		boardSrv: boardSrv,
	}
}

func (s FiberApiServer) Listen(addr string) error {
	return s.server.Listen(addr)
}

func (s FiberApiServer) Shutdown() error {
	return s.server.Shutdown()
}

func Api(
	conf *config.Config,
	lc fx.Lifecycle,
	api *FiberApiServer,
) {

	Router(api)
	SwaggerRoute(api)

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
