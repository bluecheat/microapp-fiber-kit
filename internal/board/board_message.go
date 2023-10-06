package board

type CreateBoardRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetBoardRequest struct {
	BoardId uint `params:"id"`
}

type GetBoardsRequest struct {
	Title string `query:"title"`
}

type BoardsMsg struct {
	Boards []*BoardMsg `json:"boards"`
}

type BoardMsg struct {
	BoardId   uint   `json:"boardId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Writer    string `json:"userEmail"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
