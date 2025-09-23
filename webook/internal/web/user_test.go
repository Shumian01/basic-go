package web

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"webook/internal/domain"
	"webook/internal/service"

	"testing"
	svcmocks "webook/internal/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestEncrypt(t *testing.T) {
	password := "123456"
	//加密后的数据
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	//相等返回nil
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	assert.NoError(t, err)
}
func TestNil(t *testing.T) {
	testTypeAssert(nil)

}
func testTypeAssert(c any) {
	claims := c.(*UserClaims)
	println(claims.Uid)
}

func TestUserHandler_Signup(t *testing.T) {
	testCase := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) service.UserService
		reqBody  string
		wantCode int
		wantBody string
	}{
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(nil)
				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl201515",
	"confirmPassword":"xzl201515"
}
`,
			wantCode: 200,
			wantBody: "注册成功",
		},

		{
			name: "参数不对 bind失败",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)

				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl201515"
	"confirmPassword":"xzl
}
`,
			wantCode: 400,
			wantBody: "请求格式错误",
		},

		{
			name: "邮箱格式错误",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)

				return usersvc
			},
			reqBody: `
{
	"email":"12q.com",
	"password":"xzl201515",
	"confirmPassword":"xzl201515"
}
`,
			wantCode: 400,
			wantBody: "邮箱格式不对",
		},

		{
			name: "两次输入密码不匹配",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)

				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl201515",
	"confirmPassword":"xzl25"
}
`,
			wantCode: 400,
			wantBody: "两次密码不一致",
		},

		{
			name: "密码格式不对",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl",
	"confirmPassword":"xzl"
}
`,
			wantCode: 400,
			wantBody: "密码格式不对",
		},

		{
			name: "邮箱冲突",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(service.ErrUserDuplicated)
				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl201515",
	"confirmPassword":"xzl201515"
}
`,
			wantCode: 400,
			wantBody: "邮箱冲突",
		},

		{
			name: "系统异常",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(errors.New("随便一个错误"))
				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl201515",
	"confirmPassword":"xzl201515"
}
`,
			wantCode: 200,
			wantBody: "系统异常",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//构造http请求
			server := gin.Default()
			h := NewUserHandler(tc.mock(ctrl), nil)
			h.RegisterRoutes(server)
			req, err := http.NewRequest(http.MethodPost, "/users/signup",
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
			assert.Equal(t, tc.wantBody, resp.Body.String())

		})
	}
}

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usersvc := svcmocks.NewMockUserService(ctrl)
	usersvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).
		Return(errors.New("mock error"))
	err := usersvc.SignUp(context.Background(), domain.User{
		Email: "123@qq.com",
	})
	t.Log(err)
}

func TestUserHandler_LoginJWT(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		mock     func(ctrl *gomock.Controller) service.UserService
		reqBody  string
		wantCode int
		wantBody string
	}{
		{
			name: "登录成功",
			ctx:  context.Background(),
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().Login(gomock.Any(), "123@qq.com", "xzl201515").Return(domain.User{}, nil)
				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl201515"
}
`,
			wantCode: 200,
			wantBody: "登录成功",
		},

		{
			name: "参数不对 bind失败",
			ctx:  context.Background(),
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl201
}
`,
			wantCode: 400,
			wantBody: "请求格式错误",
		},

		{
			name: "账号或密码不对",
			ctx:  context.Background(),
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().Login(gomock.Any(), "123@qq.com", "xzl201515").Return(domain.User{}, service.ErrInvalidUserOrPassword)
				return usersvc
			},
			reqBody: `
{
	"email":"123@qq.com",
	"password":"xzl201515"
}
`,
			wantCode: 200,
			wantBody: "账号或者密码不对",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//构造http请求
			server := gin.Default()
			h := NewUserHandler(tc.mock(ctrl), nil)
			h.RegisterRoutes(server)
			req, err := http.NewRequest(http.MethodPost, "/users/login",
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
			assert.Equal(t, tc.wantBody, resp.Body.String())

		})
	}
}
