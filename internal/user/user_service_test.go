package user

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"microapp-fiber-kit/domains"
	"microapp-fiber-kit/utils"
	"reflect"
	"testing"
)

// MockIBoardRepository is a mock type for the IBoardRepository
type MockIUserRepository struct {
	mock.Mock
}

func (mock MockIUserRepository) GetUser(id uint) (*domains.User, error) {
	args := mock.Called(id)
	return args.Get(0).(*domains.User), args.Error(1)
}

func (mock MockIUserRepository) GetUserByEmail(email string) (*domains.User, error) {
	args := mock.Called(email)
	return args.Get(0).(*domains.User), args.Error(1)
}

func (mock MockIUserRepository) CreateUser(user *domains.User) (*domains.User, error) {
	args := mock.Called(user)
	return args.Get(0).(*domains.User), args.Error(1)
}

var mockRepo = new(MockIUserRepository)

func TestUserService_Join(t *testing.T) {
	type fields struct {
		userRepo domains.IUserRepository
	}
	type args struct {
		req *JoinRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UserMsg
		wantErr bool
	}{
		{
			name: "회원가입 성공",
			fields: fields{
				userRepo: func() domains.IUserRepository {
					mockRepo.On("CreateUser", mock.Anything).Return(&domains.User{
						Model:    gorm.Model{ID: 1},
						Email:    "email@example.com",
						Name:     "이신일",
						Password: "1234",
					}, nil)
					return mockRepo
				}(),
			},
			args: args{
				req: &JoinRequest{
					Email:    "email@example.com",
					Name:     "이신일",
					Password: "1234",
				},
			},
			want: &UserMsg{
				UserId: 1,
				Email:  "email@example.com",
				Name:   "이신일",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				userRepo: tt.fields.userRepo,
			}
			got, err := s.Join(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Join() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Join() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	type fields struct {
		userRepo domains.IUserRepository
	}
	type args struct {
		req *LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UserMsg
		wantErr bool
	}{
		{
			name: "로그인 성공",
			fields: fields{
				userRepo: func() domains.IUserRepository {
					hashPassword, _ := utils.Hash("1234")
					mockRepo.On("GetUserByEmail", mock.Anything).Return(&domains.User{
						Model:    gorm.Model{ID: 1},
						Email:    "email@example.com",
						Name:     "이신일",
						Password: hashPassword,
					}, nil)
					return mockRepo
				}(),
			},
			args: args{
				req: &LoginRequest{
					Email:    "email@example.com",
					Password: "1234",
				},
			},
			want: &UserMsg{
				UserId: 1,
				Email:  "email@example.com",
				Name:   "이신일",
			},
			wantErr: false,
		},
		{
			name: "로그인 실패 (유저 존재하지 않음)",
			fields: fields{
				userRepo: func() domains.IUserRepository {
					mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
					return mockRepo
				}(),
			},
			args: args{
				req: &LoginRequest{
					Email:    "email@example.com",
					Password: "12345",
				},
			},
			wantErr: true,
		},
		{
			name: "로그인 실패 (패스워드 다름)",
			fields: fields{
				userRepo: func() domains.IUserRepository {
					hashPassword, _ := utils.Hash("234")
					mockRepo.On("GetUserByEmail", mock.Anything).Return(&domains.User{
						Model:    gorm.Model{ID: 1},
						Email:    "email@example.com",
						Name:     "이신일",
						Password: hashPassword,
					}, nil)
					return mockRepo
				}(),
			},
			args: args{
				req: &LoginRequest{
					Email:    "email@example.com",
					Password: "12345",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				userRepo: tt.fields.userRepo,
			}

			got, err := s.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}
