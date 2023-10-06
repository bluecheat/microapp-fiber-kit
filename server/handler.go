package server

import (
	"microapp-fiber-kit/internal/board"
)

const (
	V1 = "/v1"
)

type ServiceFunc[I interface{}, O interface{}] func(req *I) (*O, error)

func Router(
	api *FiberApiServer,
) {

	v1 := api.server.Group(V1)

	v1.Post("/board", wrapHandler[board.CreateBoardRequest, board.BoardMsg](api.boardSrv.CreateBoard))
	v1.Get("/board/:id", wrapHandler[board.GetBoardRequest, board.BoardMsg](api.boardSrv.GetBoard))
	v1.Get("/board", wrapHandler[board.GetBoardsRequest, board.BoardsMsg](api.boardSrv.GetBoards))
}
