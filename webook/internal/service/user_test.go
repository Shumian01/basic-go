package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"webook/internal/domain"
	"webook/internal/repository"
	repomocks "webook/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func Test_userService_Login(t *testing.T) {
	now := time.Now()
	testCase := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) repository.UserRepository
		ctx      context.Context
		email    string
		password string
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "登陆成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{
						Email:    "123@qq.com",
						Password: "$2a$10$1LaLx2TVxgH6xJI4Sr5LMuKJ5Luefv8oPRYkdHj84DMAiEWs8PC8O",
						Phone:    "15072128281",
						Ctime:    now,
					}, nil)
				return repo
			},
			email:    "123@qq.com",
			password: "xzl201515",
			wantUser: domain.User{
				Email:    "123@qq.com",
				Password: "$2a$10$1LaLx2TVxgH6xJI4Sr5LMuKJ5Luefv8oPRYkdHj84DMAiEWs8PC8O",
				Phone:    "15072128281",
				Ctime:    now,
			},
			wantErr: nil,
		},

		{
			name: "用户不存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{}, repository.ErrUserNotFound)
				return repo
			},
			email:    "123@qq.com",
			password: "xzl201515",
			wantUser: domain.User{},

			wantErr: ErrInvalidUserOrPassword,
		},

		{
			name: "密码不对",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{
						Email:    "123@qq.com",
						Password: "$2a$10$1LaLx2TVxgH6xJI4Sr5LMuKJ5Luefv8oPRYkdHj84DMAiEWs8PC8O",
						Phone:    "15072128281",
						Ctime:    now,
					}, nil)
				return repo
			},
			email:    "123@qq.com",
			password: "dadasdasxzl201515",
			wantUser: domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},

		{
			name: "DB错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{}, errors.New("mock db 错误"))
				return repo
			},
			email:    "123@qq.com",
			password: "xzl201515",
			wantUser: domain.User{},
			wantErr:  errors.New("mock db 错误"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			//
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := NewUserService(tc.mock(ctrl))
			u, err := svc.Login(tc.ctx, tc.email, tc.password)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, u)
		})
	}
}
func TestEncrypted(t *testing.T) {
	res, err := bcrypt.GenerateFromPassword([]byte("xzl201515"), bcrypt.DefaultCost)
	if err == nil {
		t.Log(string(res))
	}
}
