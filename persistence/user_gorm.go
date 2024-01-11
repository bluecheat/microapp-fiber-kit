package persistence

import (
	"errors"
	"microapp-fiber-kit/config/database"
	"microapp-fiber-kit/domains"
)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(database *database.Database) domains.IUserRepository {
	return &UserRepository{database: database}
}

func (r UserRepository) GetUser(id uint) (*domains.User, error) {
	user := &domains.User{}
	result := r.database.DB().First(user, "id = ?", id)
	if result.Error != nil {
		return nil, errors.New("not found user")
	}
	return user, nil
}

func (r UserRepository) GetUserByEmail(email string) (*domains.User, error) {
	user := &domains.User{}
	result := r.database.DB().First(user, "email = ?", email)
	if result.Error != nil {
		return nil, errors.New("not found user")
	}
	return user, nil
}

func (r UserRepository) CreateUser(user *domains.User) (*domains.User, error) {
	result := r.database.DB().Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
