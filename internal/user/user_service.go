package user

import (
	"errors"
	"microapp-fiber-kit/internal/domains"
	"microapp-fiber-kit/utils"
)

type UserService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Login godoc
// @Summary		로그인 API
// @Accept		json
// @Produce		json
// @Param 		LoginRequest body LoginRequest true "LoginRequest"
// @Success		200		{object}	UserMsg
// @Failure	409  		{object}  	server.Error
// @Router			/v1/login [post]
func (s UserService) Login(req *LoginRequest) (*UserMsg, error) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("이메일 혹은 비밀번호가 틀립니다")
	}

	if err := utils.VerifyHash(user.Password, req.Password); err != nil {
		return nil, errors.New("이메일 혹은 비밀번호가 틀립니다")
	}
	return &UserMsg{
		UserId: user.ID,
		Email:  user.Email,
		Name:   user.Name,
	}, nil
}

// Join godoc
// @Summary		회원가입 API
// @Accept		json
// @Produce		json
// @Param 		JoinRequest body JoinRequest true "JoinRequest"
// @Success		200		{object}	UserMsg
// @Failure	409  		{object}  	server.Error
// @Router			/v1/join [post]
func (s UserService) Join(req *JoinRequest) (*UserMsg, error) {
	hashedPassword, _ := utils.Hash(req.Password)
	user, err := s.userRepo.CreateUser(&domains.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	return &UserMsg{
		UserId: user.ID,
		Email:  user.Email,
		Name:   user.Name,
	}, nil
}
