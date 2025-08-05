package middleware

import (
	"basic-go/webook/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

// LoginJWTMiddlewareBuilder  JWT登录校验
type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}
func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//不需要登录校验的
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		//我现在用JWT校验
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			//没登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			//没登录 有人瞎搞
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		claims := &web.UserClaims{}
		//ParseWithClaims一定要传指针
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return []byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX"), nil
		})
		if err != nil {
			//没登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//err 为nil token不为nil
		if token == nil || !token.Valid || claims.Uid == 0 {
			//没登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//刷新登陆状态
		//每十秒刷新
		if claims.UserAgent != ctx.Request.UserAgent() {
			//严重的安全问题
			//要加监控
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		now := time.Now()
		if claims.ExpiresAt.Sub(now) < time.Second*50 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenStr, err := token.SignedString([]byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX"))
			if err != nil {
				//记录日志
				log.Println("jwt续约失败", err)
			}
			//产生新token
			ctx.Header("x-jwt-token", tokenStr)
		}

		ctx.Set("claims", claims)
	}
}
