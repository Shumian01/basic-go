package repository

import (
	"database/sql"
	"time"
	"webook/internal/repository/cache"
	"webook/internal/repository/dao"
)

import (
	"context"
	"webook/internal/domain"
)

var (
	ErrUserDuplicated = dao.ErrUserDuplicated
	ErrUserNotFound   = dao.ErrUserNotFound
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
func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	u, err := r.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return r.entityToDomain(u), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return r.entityToDomain(u), nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, r.domainToEntity(u))
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
	u = r.entityToDomain(ue)
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
func (r *UserRepository) domainToEntity(u domain.User) dao.User {
	return dao.User{
		Id: u.Id,
		Email: sql.NullString{
			String: u.Email,
			Valid:  u.Email != "",
		},
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		Password: u.Password,
		Ctime:    u.Ctime.UnixMilli(),
	}
}

func (r *UserRepository) entityToDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email.String,
		Phone:    u.Phone.String,
		Password: u.Password,
		Ctime:    time.UnixMilli(u.Ctime),
	}
}
