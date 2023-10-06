package board

import (
	"errors"
	"gorm.io/gorm"
	"microapp-fiber-kit/database"
	"microapp-fiber-kit/internal/domains"
)

type IBoardRepository interface {
	GetBoard(id uint) (*domains.Board, error)
	SearchBoard(search map[string]string) ([]*domains.Board, error)
	CreateBoard(board *domains.Board) (*domains.Board, error)
}

type BoardRepository struct {
	database *database.Database
}

func NewBoardRepository(database *database.Database) IBoardRepository {
	return &BoardRepository{database: database}
}

func (r BoardRepository) GetBoard(id uint) (*domains.Board, error) {
	board := &domains.Board{}
	result := r.database.DB().First(board, "id = ?", id)
	if result.Error != nil {
		return nil, errors.New("not found board")
	}
	return board, nil
}

func (r BoardRepository) SearchBoard(search map[string]string) ([]*domains.Board, error) {
	var boards []*domains.Board
	sql := r.database.DB()
	title, ok := search["title"]
	if ok && title != "" {
		sql = sql.Scopes(likeTitle(title))
	}
	sql.Find(&boards)
	return boards, nil
}

func (r BoardRepository) CreateBoard(board *domains.Board) (*domains.Board, error) {
	result := r.database.DB().Create(board)
	if result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}

func likeTitle(title string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("title LIKE ?", "%"+title+"%")
	}
}
