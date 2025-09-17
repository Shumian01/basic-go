package main

func main() {

	server := InitWebServer()
	server.Run(":8080")
}

//func InitUserWeb() *gin.Engine {
//	server := gin.Default()
//	//redisClient := redis.NewClient(&redsi)
//	//server.Use(ratelimit.NewBuilder())
//	//redisClient := redis.NewClient(&redis.Options{
//	//	Addr: config.Config.Redis.Addr,
//	//})
//	//server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())
//
//	server.Use(cors.New(cors.Config{
//		AllowOrigins: []string{"http://localhost:3000"},
//		//AllowMethods: []string{"POST", "GET"}, //不写等于都支持
//		AllowHeaders: []string{"Content-type", "Authorization"}, //
//		//不加这个前端拿不到
//		ExposeHeaders: []string{"x-jwt-token"},
//		//是否允许cookie之类的东西
//		AllowCredentials: true,
//		AllowOriginFunc: func(origin string) bool {
//			if strings.HasPrefix(origin, "http://localhost") {
//				//你的开发环境
//				return true
//			}
//			return strings.Contains(origin, "yourcompany.com")
//		},
//		MaxAge: 12 * time.Hour,
//	}))
//	//1.cookie装session
//	//2.memstore 使用内存
//	//3.使用Redis
//	//store := cookie.NewStore([]byte("secret"))
//	//
//	//第一个参数 最大空闲连接数量
//	//2 tcp
//	//3 4 连接信息和密码
//	//5 key
//	//store, err := redis.NewStore(32, "tcp", "localhost:6379", "", "",
//	//	[]byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX"),
//	//	[]byte("iIUJ20V9jJlEYjlfkf17Rk8deT2v2Qo7"))
//	//if err != nil {
//	//	panic(err)
//	//}
//	//store := memstore.NewStore([]byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX"), []byte("iIUJ20V9jJlEYjlfkf17Rk8deT2v2Qo7"))
//	//server.Use(sessions.Sessions("mysession", store))
//
//	//server.Use(middleware.NewLoginMiddlewareBuilder().IgnorePaths("/users/login").IgnorePaths("/users/signup").Build())
//	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
//		IgnorePaths("/users/signup").
//		IgnorePaths("/users/login_sms/code/send").
//		IgnorePaths("/users/login_sms").
//		IgnorePaths("/users/login").Build())
//	return server
//}
//
//func InitUser(db *gorm.DB, rdb redis.Cmdable) *web.UserHandler {
//	ud := dao.NewUserDAO(db)
//	repo := repository.NewUserRepository(ud, nil)
//	svc := service.NewUserService(repo)
//	codeCache := cache.NewCodeCache(rdb)
//	codeRepo := repository.NewCodeRepository(codeCache)
//	smsSvc := memory.NewService()
//	codeSvc := service.NewCodeService(codeRepo, smsSvc)
//	u := web.NewUserHandler(svc, codeSvc)
//	return u
//}
