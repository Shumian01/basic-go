package web

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
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
