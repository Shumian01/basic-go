package ratelimit

import "context"

type Limiter interface {
	//Limite 有没有触发限流,key 就是限流对象
	//bool true代表限流
	//error 限流器有没有错误
	Limite(ctx context.Context, key string) (bool, error)
}
