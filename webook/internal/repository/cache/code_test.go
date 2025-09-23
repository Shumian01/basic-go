package cache

import (
	"context"
	"errors"
	"testing"
	"webook/internal/repository/cache/redismocks"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRedisCodeCache_Set(t *testing.T) {
	testCases := []struct {
		name string

		mock func(ctrl *gomock.Controller) redis.Cmdable
		//输入
		ctx   context.Context
		biz   string
		phone string
		code  string
		//输出
		wantErr error
	}{
		{
			name: "验证码设置成功",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetErr(nil)
				res.SetVal(int64(0))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode,
					[]string{"phone_code:login:150"}, []any{"123456"}).Return(res)
				return cmd
			},
			ctx:     context.Background(),
			biz:     "login",
			phone:   "150",
			code:    "123456",
			wantErr: nil,
		},

		{
			name: "redis错误",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetErr(errors.New("mock redis 错误"))
				res.SetVal(int64(0))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode,
					[]string{"phone_code:login:150"}, []any{"123456"}).Return(res)
				return cmd
			},
			ctx:     context.Background(),
			biz:     "login",
			phone:   "150",
			code:    "123456",
			wantErr: errors.New("mock redis 错误"),
		},

		{
			name: "发生太频繁",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetErr(nil)
				res.SetVal(int64(-1))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode,
					[]string{"phone_code:login:150"}, []any{"123456"}).Return(res)
				return cmd
			},
			ctx:     context.Background(),
			biz:     "login",
			phone:   "150",
			code:    "123456",
			wantErr: ErrSetCodeTooMany,
		},

		{
			name: "系统错误",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetErr(nil)
				res.SetVal(int64(-10))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode,
					[]string{"phone_code:login:150"}, []any{"123456"}).Return(res)
				return cmd
			},
			ctx:     context.Background(),
			biz:     "login",
			phone:   "150",
			code:    "123456",
			wantErr: errors.New("系统错误"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cache := NewCodeCache(tc.mock(ctrl))
			err := cache.Set(tc.ctx, tc.biz, tc.phone, tc.code)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
