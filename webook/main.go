package main

import (
	"basic-go/webook/internal/repository"
	"basic-go/webook/internal/repository/dao"
	"basic-go/webook/internal/service"
	"basic-go/webook/internal/web"
	"basic-go/webook/internal/web/middleware"
	"basic-go/webook/pkg/ginx/middleware/ratelimit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func main() {
	//db := InitDB()          //初始化DB
	//server := InitUserWeb() //初始化server
	//u := InitUser(db)
	//u.RegisterUser(server)
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello,xzl")
	})
	server.Run(":8080")
}

func InitUserWeb() *gin.Engine {
	server := gin.Default()
	//redisClient := redis.NewClient(&redsi)
	//server.Use(ratelimit.NewBuilder())
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		//AllowMethods: []string{"POST", "GET"}, //不写等于都支持
		AllowHeaders: []string{"Content-type", "Authorization"}, //
		//不加这个前端拿不到
		ExposeHeaders: []string{"x-jwt-token"},
		//是否允许cookie之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//你的开发环境
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	//1.cookie装session
	//2.memstore 使用内存
	//3.使用Redis
	//store := cookie.NewStore([]byte("secret"))
	//
	//第一个参数 最大空闲连接数量
	//2 tcp
	//3 4 连接信息和密码
	//5 key
	//store, err := redis.NewStore(32, "tcp", "localhost:6379", "", "",
	//	[]byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX"),
	//	[]byte("iIUJ20V9jJlEYjlfkf17Rk8deT2v2Qo7"))
	//if err != nil {
	//	panic(err)
	//}
	store := memstore.NewStore([]byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX"), []byte("iIUJ20V9jJlEYjlfkf17Rk8deT2v2Qo7"))
	server.Use(sessions.Sessions("mysession", store))

	//server.Use(middleware.NewLoginMiddlewareBuilder().IgnorePaths("/users/login").IgnorePaths("/users/signup").Build())
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").Build())
	return server
}

func InitUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		//我只会在初始化过程panic
		//panic 相当于整个goroutine 结束
		//一旦初始化过程出错 应用就不要启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)

	}
	return db
}
