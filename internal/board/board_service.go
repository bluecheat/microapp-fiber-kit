package board

import (
	"microapp-fiber-kit/internal/domains"
	"microapp-fiber-kit/utils"
)

type BoardService struct {
	boardRepo IBoardRepository
}

func NewBoardService(boardRepo IBoardRepository) *BoardService {
	return &BoardService{boardRepo: boardRepo}
}

// CreateBoard godoc
// @Summary		게시판 등록 API
// @Accept		json
// @Produce		json
// @Param 		board.CreateBoardRequest body board.CreateBoardRequest true "CreateBoardRequest"
// @Success		200		{object}	board.BoardMsg
// @Failure	409  		{object}  	server.Error
// @Router			/v1/board [post]
func (s *BoardService) CreateBoard(req *CreateBoardRequest) (*BoardMsg, error) {
	newBoard := &domains.Board{
		Title:   req.Title,
		Content: req.Content,
		Writer:  "user1@email.com",
	}
	newBoard, err := s.boardRepo.CreateBoard(newBoard)
	if err != nil {
		return nil, err
	}
	createTime, _ := utils.ParseTime(newBoard.Model, "2006-01-02 15:04:05")
	return &BoardMsg{
		BoardId:   newBoard.ID,
		Title:     newBoard.Title,
		Content:   newBoard.Content,
		CreatedAt: createTime,
		Writer:    newBoard.Writer,
	}, nil
}

// GetBoard godoc
// @Summary		게시판 조회 API
// @Accept		json
// @Produce		json
// @Param			id	path		int	true	"Board ID"
// @Success		200		{object}	board.BoardMsg
// @Failure	409  		{object}  	server.Error
// @Router			/v1/board/{id} [get]
func (s *BoardService) GetBoard(req *GetBoardRequest) (*BoardMsg, error) {
	board, err := s.boardRepo.GetBoard(req.BoardId)
	if err != nil {
		return nil, err
	}
	createTime, updateTime := utils.ParseTime(board.Model, "2006-01-02 15:04:05")
	return &BoardMsg{
		BoardId:   board.ID,
		Title:     board.Title,
		Content:   board.Content,
		CreatedAt: createTime,
		UpdatedAt: updateTime,
		Writer:    board.Writer,
	}, nil
}

// GetBoards godoc
// @Summary		게시판 목록 조회 API
// @Accept		json
// @Produce		json
// @Param 		board.GetBoardsRequest query board.GetBoardsRequest true "GetBoardsRequest"
// @Success		200		{object}	board.BoardsMsg
// @Failure	409  		{object}  	server.Error
// @Router			/v1/board [get]
func (s *BoardService) GetBoards(req *GetBoardsRequest) (*BoardsMsg, error) {
	var boards []*domains.Board

	boards, err := s.boardRepo.SearchBoard(map[string]string{"title": req.Title})
	if err != nil {
		return nil, err
	}
	boardsMsg := make([]*BoardMsg, len(boards))
	for i, board := range boards {
		createTime, updateTime := utils.ParseTime(board.Model, "2006-01-02 15:04:05")
		boardsMsg[i] = &BoardMsg{
			BoardId:   board.ID,
			Title:     board.Title,
			Content:   board.Content,
			CreatedAt: createTime,
			UpdatedAt: updateTime,
			Writer:    board.Writer,
		}
	}
	return &BoardsMsg{
		Boards: boardsMsg,
	}, nil
}
