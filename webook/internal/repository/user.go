package repository

import (
	"basic-go/webook/internal/repository/dao"
)

import (
	"basic-go/webook/internal/domain"
	"context"
)

var (
	ErrUserDuplicatedEmail = dao.ErrUserDuplicatedEmail
	ErrUserNotFound        = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
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

func (r *UserRepository) FindById(int64) {
	//先从cache中找
	//再从dao里面找
	//找到了回写cache
}
