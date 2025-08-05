package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicatedEmail = errors.New("邮箱冲突")
	ErrUserNotFound        = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}
func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?").Error
	return u, err
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	//存毫秒
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlerr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo = 1062
		if mysqlerr.Number == 1062 {
			//邮箱冲突
			return ErrUserDuplicatedEmail
		}
	}
	return err
}

// User 直接对应数据库表结构
// 有些人叫做entity  model
type User struct {
	Id int64 `gorm:"primary_key,auto_increment"`
	//全部用户唯一   加上唯一索引
	Email    string `gorm:"unique"`
	Password string

	//这里添加新字段

	//创建时间 毫秒
	Ctime int64
	//更新时间 毫秒
	Utime int64
}
