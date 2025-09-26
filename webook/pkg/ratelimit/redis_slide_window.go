package ratelimit

import (
	"context"
	_ "embed"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:embed slide_window.lua
var luaSlideWidow string

//Redis上滑动窗口算法限流器实现

type RedisSlidingWindowLimiter struct {
	cmd redis.Cmdable
	//窗口大小
	interval time.Duration
	// 阈值
	rate int
	//interval内允许 rate个请求
	//1s 内允许 3000个请求
}

func NewRedisSlidingWindowLimiter(cmd redis.Cmdable,
	interval time.Duration, rate int) Limiter {
	return &RedisSlidingWindowLimiter{
		cmd:      cmd,
		interval: interval,
		rate:     rate,
	}
}

func (r RedisSlidingWindowLimiter) Limite(ctx context.Context, key string) (bool, error) {
	return r.cmd.Eval(ctx, luaSlideWidow, []string{key},
		r.interval.Milliseconds(), r.rate, time.Now().UnixMilli()).Bool()
}
