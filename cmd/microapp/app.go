package main

import (
	"go.uber.org/fx"
	"microapp-fiber-kit/config"
	"microapp-fiber-kit/database"
	"microapp-fiber-kit/internal/board"
	"microapp-fiber-kit/server"
)

//type Application interface {
//	Run() error
//	Stop() error
//}

func MicroApp(conf *config.Config) *fx.App {
	return fx.New(
		fx.Provide(
			append(
				providers(),
				func() *config.Config { return conf },
			)...,
		),
		fx.Invoke(
			invokers()...,
		),
	)
}

// 공급
func providers() []interface{} {
	return []interface{}{
		database.NewDatabase,
		board.NewBoardRepository,
		board.NewBoardService,
		server.NewFiberApiServer,
	}
}

// 호출
func invokers() []interface{} {
	return []interface{}{
		database.AutoMigration,
		server.Api,
	}
}
