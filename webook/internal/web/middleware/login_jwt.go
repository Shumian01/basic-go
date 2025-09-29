package middleware

import (
	"net/http"
	ijwt "webook/internal/web/jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// LoginJWTMiddlewareBuilder  JWT登录校验
type LoginJWTMiddlewareBuilder struct {
	paths []string
	ijwt.Handler
}

func NewLoginJWTMiddlewareBuilder(jwtHdl ijwt.Handler) *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{
		Handler: jwtHdl,
	}
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
		tokenStr := l.ExtractToken(ctx)
		claims := &ijwt.UserClaims{}
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
		if claims.UserAgent != ctx.Request.UserAgent() {
			//严重的安全问题
			//要加监控
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//验证ssid
		// if redis 崩了 return
		err = l.CheckSession(ctx, claims.Ssid)
		if err != nil {
			// 要么redis出问题 要么已经退出登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("user", claims)
	}
}
