//go:build wireinject

package integration

import (
	"webook/internal/repository"
	"webook/internal/repository/cache"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/internal/web"
	"webook/ioc"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		//最基础第三方依赖
		ioc.InitDB, ioc.InitRedis,
		//初始化DAO
		dao.NewUserDAO,
		//初始化cache
		cache.NewUserCache, cache.NewCodeCache,
		repository.NewUserRepository, repository.NewCodeRepository,
		service.NewUserService, service.NewCodeService,
		//基于内存实现
		ioc.InitSMSService,
		web.NewUserHandler,
		//中间件呢?
		//注册路由呢
		//gin.Default,
		ioc.WebServer,
		ioc.InitMiddlewares,
	)
	return new(gin.Engine)
}
