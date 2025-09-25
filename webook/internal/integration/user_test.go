package integration

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"webook/internal/web"
	"webook/ioc"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_e2e_SendLoginSMSCode(t *testing.T) {
	server := InitWebServer()
	rdb := ioc.InitRedis()
	testCases := []struct {
		name string
		//你要考虑准备数据 以及验证数据
		before func(t *testing.T)
		//以及验证数据 数据库的数据 对不对 redis数据对不对
		after    func(t *testing.T)
		reqBody  string
		wantCode int
		wantBody web.Result
	}{
		{
			name: "发送成功",
			before: func(t *testing.T) {
				//不需要 也就是redis 什么数据也没有
			},
			after: func(t *testing.T) {
				//超时时间2
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				//清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:15072128281").Result()
				cancel()
				assert.NoError(t, err)
				assert.True(t, len(val) == 6)
			},
			reqBody: `
{
	"phone":"15072128281"
}
`,
			wantCode: http.StatusOK,
			wantBody: web.Result{
				Msg: "发送成功",
			},
		},

		{
			name: "发送太频繁",
			before: func(t *testing.T) {
				//这个手机号已经有一个验证码
				ctx, cancel := context.WithTimeout(context.Background(),
					3*time.Second)
				//清理数据
				_, err := rdb.Set(ctx, "phone_code:login:15072128281", "123456",
					9*time.Minute+30*time.Second).Result()
				cancel()
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//超时时间2
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				//清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:15072128281").Result()
				cancel()
				assert.NoError(t, err)
				assert.Equal(t, "123456", val)
			},
			reqBody: `
{
	"phone":"15072128281"
}
`,
			wantCode: http.StatusOK,
			wantBody: web.Result{
				Msg: "发送太频繁,请稍后再试",
			},
		},

		{
			name: "系统错误",
			before: func(t *testing.T) {
				//这个手机号已经有一个验证码,但是没有过期时间 ->系统错误
				ctx, cancel := context.WithTimeout(context.Background(),
					3*time.Second)
				//清理数据
				_, err := rdb.Set(ctx, "phone_code:login:15072128281", "123456", 0).Result()
				cancel()
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//超时时间2
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				//清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:15072128281").Result()
				cancel()
				assert.NoError(t, err)
				assert.Equal(t, "123456", val)
			},
			reqBody: `
{
	"phone":"15072128281"
}
`,
			wantCode: http.StatusOK,
			wantBody: web.Result{
				Code: 5,
				Msg:  "系统错误",
			},
		},

		{
			name: "手机号码为空",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {

			},
			reqBody: `
{
	"phone":""
}
`,
			wantCode: http.StatusOK,
			wantBody: web.Result{
				Code: 4,
				Msg:  "输入有误",
			},
		},

		{
			name: "数据格式错误",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {

			},
			reqBody: `
{
	"phone":"
}
`,
			wantCode: 400,
			wantBody: web.Result{
				Code: 4,
				Msg:  "输入有误",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//构造http请求
			tc.before(t)
			req, err := http.NewRequest(http.MethodPost, "/users/login_sms/code/send",
				bytes.NewBuffer([]byte(tc.reqBody)))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			//这里继续使用req
			resp := httptest.NewRecorder()
			t.Log(resp)
			//Http请求进去gin框架的入口
			//当你这样调用的时候 gin就会处理这个请求
			// 响应回写到resp
			server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var webRes web.Result
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			require.NoError(t, err)
			assert.Equal(t, tc.wantBody, webRes)
			tc.after(t)
		})
	}
}
