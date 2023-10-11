package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"gopkg.in/natefinch/lumberjack.v2"
	"microapp-fiber-kit/config"
	"strings"
)

type FiberApiServer struct {
	conf   *config.Config
	Server *fiber.App
}

func initializeServer(conf *config.Config) *fiber.App {
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

	return server
}

func NewFiberApiServer(
	conf *config.Config,
) *FiberApiServer {
	server := initializeServer(conf)
	return &FiberApiServer{
		Server: server,
	}
}

func (s FiberApiServer) Listen(addr string) error {
	return s.Server.Listen(addr)
}

func (s FiberApiServer) Shutdown() error {
	return s.Server.Shutdown()
}
