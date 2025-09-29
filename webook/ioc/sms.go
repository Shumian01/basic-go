package ioc

import (
	"webook/internal/service/sms"
	"webook/internal/service/sms/memory"

	"github.com/redis/go-redis/v9"
)

// 装饰器
func InitSMSService(cmd redis.Cmdable) sms.Service {
	//svc := ratelimit.NewRatelimitSMSService(memory.NewService(),
	//	ratelimit2.NewRedisSlidingWindowLimiter(cmd, time.Second, 100))

	return memory.NewService()
}
