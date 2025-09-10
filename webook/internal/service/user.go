package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"time"
	"webook/internal/domain"
	"webook/internal/repository"
)

var ErrUserDuplicatedEmail = repository.ErrUserDuplicatedEmail
var ErrInvalidUserOrPassword = errors.New("账号或者密码不对")
var ErrUserNotFound = repository.ErrUserNotFound

type UserService struct {
	repo  *repository.UserRepository
	redis *redis.Client
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	err = svc.repo.Create(ctx, u)
	if err != nil {
		return err
	}
	//
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	//要求id不为0
	err = svc.redis.Set(ctx, fmt.Sprintf("user:info:%d", u.Id), val, time.Minute*60).Err()
	return err
}

func (svc *UserService) Login(ctx context.Context, email string, password string) (domain.User, error) {

	//先找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	//比较密码

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		//DEBUG
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

//func (svc *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
//	u, err := svc.repo.FindById(ctx, id)
//	//没这个数据 去数据库找
//	return do
//}
