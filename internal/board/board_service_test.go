package board

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"microapp-fiber-kit/internal/domains"
	"testing"
)

// MockIBoardRepository is a mock type for the IBoardRepository
type MockIBoardRepository struct {
	mock.Mock
}

func (mock *MockIBoardRepository) CreateBoard(board *domains.Board) (*domains.Board, error) {
	args := mock.Called(board)
	return args.Get(0).(*domains.Board), args.Error(1)
}

func (mock *MockIBoardRepository) GetBoard(id uint) (*domains.Board, error) {
	args := mock.Called(id)
	return args.Get(0).(*domains.Board), args.Error(1)
}

func (mock *MockIBoardRepository) SearchBoard(criteria map[string]string) ([]*domains.Board, error) {
	args := mock.Called(criteria)
	return args.Get(0).([]*domains.Board), args.Error(1)
}

var mockRepo = new(MockIBoardRepository)

// TestCreateBoard tests the CreateBoard method
func TestCreateBoard(t *testing.T) {
	board := &domains.Board{
		Title:   "Test Title",
		Content: "Test Content",
		Writer:  "user1@email.com",
	}
	mockRepo.On("CreateBoard", mock.Anything).Return(board, nil)

	service := NewBoardService(mockRepo)
	request := &CreateBoardRequest{
		Title:   "Test Title",
		Content: "Test Content",
	}

	response, err := service.CreateBoard(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, board.Title, response.Title)
	assert.Equal(t, board.Content, response.Content)
	assert.Equal(t, board.Writer, response.Writer)

	mockRepo.AssertExpectations(t)
}

// TestGetBoard tests the GetBoard method
func TestGetBoard(t *testing.T) {
	// Similar to TestCreateBoard, implement the test for GetBoard
	service := NewBoardService(mockRepo)
	board := &domains.Board{
		Model: gorm.Model{
			ID: 1,
		},
		Title:   "Test Title",
		Content: "Test Content",
		Writer:  "user1@email.com",
	}
	mockRepo.On("GetBoard", uint(1)).Return(board, nil)

	request := &GetBoardRequest{BoardId: 1}
	response, err := service.GetBoard(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, board.ID, response.BoardId)
	assert.Equal(t, board.Title, response.Title)
	assert.Equal(t, board.Content, response.Content)
	assert.Equal(t, board.Writer, response.Writer)
}

// TestGetBoards tests the GetBoards method
func TestGetBoards(t *testing.T) {
	mockRepo := new(MockIBoardRepository)
	boards := []*domains.Board{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Title:   "Test Title 1",
			Content: "Test Content 1",
			Writer:  "user1@email.com",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Title:   "Test Title 2",
			Content: "Test Content 2",
			Writer:  "user2@email.com",
		},
	}

	mockRepo.On("SearchBoard", mock.Anything).Return(boards, nil)
	service := NewBoardService(mockRepo)

	request := &GetBoardsRequest{Title: "Test"}
	response, err := service.GetBoards(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.Boards, len(boards))
	for i, board := range response.Boards {
		assert.Equal(t, boards[i].ID, board.BoardId)
		assert.Equal(t, boards[i].Title, board.Title)
		assert.Equal(t, boards[i].Content, board.Content)
		assert.Equal(t, boards[i].Writer, board.Writer)
	}
}
