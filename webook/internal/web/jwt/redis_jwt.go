package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	AtKey = []byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX")
	RtKey = []byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PF0")
)

type RedisJwtHandler struct {
	//assess_token key , refresh_token key
	cmd redis.Cmdable
}

func NewRedisJWTHandler(cmd redis.Cmdable) Handler {
	return &RedisJwtHandler{
		cmd: cmd,
	}

}

func (h *RedisJwtHandler) SetRefreshToken(ctx *gin.Context, uid int64, ssid string) error {
	claims := RefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
		Uid:  uid,
		Ssid: ssid,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	refreshTokenStr, err := refreshToken.SignedString(RtKey)
	if err != nil {
		return err
	}
	ctx.Header("x-refresh-token", refreshTokenStr)
	return err
}

func (h *RedisJwtHandler) ExtractToken(ctx *gin.Context) string {
	tokenHeader := ctx.GetHeader("Authorization")
	segs := strings.SplitN(tokenHeader, " ", 2)
	if len(segs) != 2 {
		return ""
	}
	return segs[1]
}

func (h *RedisJwtHandler) SetLoginToken(ctx *gin.Context, uid int64) error {
	ssid := uuid.New().String()
	err := h.SetJWTToken(ctx, uid, ssid)
	if err != nil {
		return err
	}
	err = h.SetRefreshToken(ctx, uid, ssid)
	if err != nil {
		return err
	}
	return nil
}

func (h *RedisJwtHandler) ClearToken(ctx *gin.Context) error {
	//token换成非法值
	ctx.Header("x-jwt-token", "")
	ctx.Header("x-refresh-token", "")
	uc := ctx.MustGet("user").(*UserClaims)
	err := h.cmd.Set(ctx, fmt.Sprintf("user:ssid:%s", uc.Ssid), "", time.Hour*24*7).Err()
	return err
}

func (h *RedisJwtHandler) CheckSession(ctx *gin.Context, ssid string) error {
	_, err := h.cmd.Exists(ctx, fmt.Sprintf("user:ssid:%s", ssid)).Result()
	return err
}

func (h *RedisJwtHandler) SetJWTToken(ctx *gin.Context, uid int64, ssid string) error {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
		Uid:       uid,
		UserAgent: ctx.Request.UserAgent(),
		Ssid:      ssid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(AtKey)
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return err
}
