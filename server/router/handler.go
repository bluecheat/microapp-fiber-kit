package router

import (
	"github.com/gofiber/fiber/v2"
	"microapp-fiber-kit/internal/board"
	"microapp-fiber-kit/internal/user"
)

const (
	V1 = "/v1"
)

func Route(
	router fiber.Router,
	boardSrv *board.BoardService,
	userSrv *user.UserService,
) {
	v1 := router.Group(V1)

	v1.Post("/board", wrapHandler[board.CreateBoardRequest, board.BoardMsg](boardSrv.CreateBoard))
	v1.Get("/board/:id", wrapHandler[board.GetBoardRequest, board.BoardMsg](boardSrv.GetBoard))
	v1.Get("/board", wrapHandler[board.GetBoardsRequest, board.BoardsMsg](boardSrv.GetBoards))

	v1.Post("/login", wrapHandler[user.LoginRequest, user.UserMsg](userSrv.Login))
	v1.Post("/join", wrapHandler[user.JoinRequest, user.UserMsg](userSrv.Join))
}
