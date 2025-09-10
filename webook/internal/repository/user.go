package repository

import (
	"webook/internal/repository/cache"
	"webook/internal/repository/dao"
)

import (
	"context"
	"webook/internal/domain"
)

var (
	ErrUserDuplicatedEmail = dao.ErrUserDuplicatedEmail
	ErrUserNotFound        = dao.ErrUserNotFound
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDAO, cache *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: cache,
	}
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Email:    u.Email,
		Password: u.Password,
		Id:       u.Id,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
	//操作缓存
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	//先从cache中找
	//再从dao里面找
	//找到了回写cache
	u, err := r.cache.Get(ctx, id)
	if err == nil {
		//必然有数据
		return u, nil
	}

	//if err == lua.ErrUserNotFound {
	//	//没这个数据 去数据库找
	//}
	ue, err := r.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	u = domain.User{
		Id:       ue.Id,
		Email:    ue.Email,
		Password: ue.Password,
	}
	err = r.cache.Set(ctx, id, u)
	go func() {
		if err != nil {
			//缓存失败
			//打个日志 监控

		}
	}()
	return u, err

	//这里这么办？要不要在数据库加载？
	//选加载 ---做好兜底,万一redis真的挂了 要保护数据库
	//数据库限流保护数据库

	//选不加载 ---用户体验差一点

	//缓存里面有数据
	//缓存没有数据
	//缓存出错了
}
