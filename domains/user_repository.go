package domains

type IUserRepository interface {
	GetUser(id uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) (*User, error)
}
