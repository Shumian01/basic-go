package web

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
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
		name string
	}{}

	//构造http请求
	req, err := http.NewRequest(http.MethodPost, "/users/signup",
		bytes.NewBuffer([]byte(`
{	
	"email":"123@qq.com",
	"password":"123456"
}
`)))
	require.NoError(t, err)
	//这里你就可以继续使用 req
	resp := httptest.NewRecorder()

	h := NewUserHandler()
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			//我在这里怎么拿到响应
			//headler := NewUserHandler(nil, nil)
			//ctx := &gin.Context{}
			//headler.Signup(ctx)
		})
	}
}
