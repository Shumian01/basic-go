package cache

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"time"
	"webook/internal/domain"
)

var ErrUserNotFound = redis.Nil

type UserCache interface {
	Get(ctx context.Context, id int64) (domain.User, error)
	Set(ctx context.Context, id int64, u domain.User) error
}
type RedisUserCache struct {
	//传单机redis可以
	//传 cluster 的redis也可以
	client     redis.Cmdable
	expiration time.Duration
}

//A 用到了B, B一定是接口
//A 用到了B ,B一定是A的字段
//A 用到了B, A绝对不初始化B 而是从外面注入
func NewUserCache(client redis.Cmdable) UserCache {
	return &RedisUserCache{
		client:     client,
		expiration: time.Minute * 15,
	}
}

//只要err 为nil 就认为缓存里有数据
//如果没有数据 返回一个特定的error
func (cache *RedisUserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := cache.Key(id)
	//数据不存在 err=redis.Nil
	val, err := cache.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}

	var u domain.User
	err = json.Unmarshal(val, &u)

	return u, err
}

func (cache *RedisUserCache) Set(ctx context.Context, id int64, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := cache.Key(id)
	err = cache.client.Set(ctx, key, val, cache.expiration).Err()
	return err
}
func (cache *RedisUserCache) Key(id int64) string {
	//user:info:123
	return fmt.Sprintf("user:info:%d", id)
}
