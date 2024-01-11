package domains

type IBoardRepository interface {
	GetBoard(id uint) (*Board, error)
	SearchBoard(search map[string]string) ([]*Board, error)
	CreateBoard(board *Board) (*Board, error)
}
