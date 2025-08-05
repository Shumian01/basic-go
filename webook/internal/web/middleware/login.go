package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}
func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

//func CheckLogin() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		for _, path := range IgnorePaths {
//			if ctx.Request.URL.Path == path {
//				return
//			}
//		}
//		sess := sessions.Default(ctx)
//		id := sess.Get("userId")
//		if id == nil {
//			//没有登录
//			ctx.AbortWithStatus(http.StatusUnauthorized)
//			return
//		}
//	}
//}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		//不需要登录校验的
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		if id == nil {
			//没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//刷新登录状态
		//我在某个地方
		updateTime := sess.Get("update_time") //上一次更新时间
		sess.Set("userId", id)
		sess.Options(sessions.Options{
			MaxAge: 60,
		})
		now := time.Now().UnixMilli()
		//刚登录还没刷新
		if updateTime != nil {
			sess.Set("update_time", now)

			sess.Save()
			return
		}
		//updata是有的
		updateTimeVal, _ := updateTime.(int64)

		if now-updateTimeVal > 1000*60 {
			sess.Set("update_time", now)
			sess.Save()
		}
	}
}
