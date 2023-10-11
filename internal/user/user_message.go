package user

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type JoinRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserMsg struct {
	UserId uint   `json:"userId"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}
