package router

import (
	"microapp-fiber-kit/internal/board"
	"microapp-fiber-kit/server"
)

const (
	V1 = "/v1"
)

func Router(
	api *server.FiberApiServer,
) {

	v1 := api.Server.Group(V1)

	v1.Post("/board", wrapHandler[board.CreateBoardRequest, board.BoardMsg](api.BoardSrv.CreateBoard))
	v1.Get("/board/:id", wrapHandler[board.GetBoardRequest, board.BoardMsg](api.BoardSrv.GetBoard))
	v1.Get("/board", wrapHandler[board.GetBoardsRequest, board.BoardsMsg](api.BoardSrv.GetBoards))
}
