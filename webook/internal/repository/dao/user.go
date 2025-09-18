package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserNotFound   = gorm.ErrRecordNotFound
	ErrUserDuplicated = errors.New("邮箱or号码冲突")
)

type UserDAO interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
	FindById(ctx context.Context, id int64) (User, error)
	Insert(ctx context.Context, u User) error
}
type GORMUserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}
func (dao *GORMUserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?").Error
	return u, err
}
func (dao *GORMUserDAO) FindByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?").Error
	return u, err
}
func (dao *GORMUserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?").Error
	return u, err
}

func (dao *GORMUserDAO) Insert(ctx context.Context, u User) error {
	//存毫秒
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlerr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo = 1062
		if mysqlerr.Number == 1062 {
			//邮箱冲突
			return ErrUserDuplicated
		}
	}
	return err
}

// User 直接对应数据库表结构
// 有些人叫做entity  model
type User struct {
	Id int64 `gorm:"primary_key,auto_increment"`
	//全部用户唯一   加上唯一索引
	Email    sql.NullString `gorm:"unique"`
	Password string
	//唯一索引允许有多个空值
	Phone sql.NullString `gorm:"unique"`
	//这里添加新字段

	//创建时间 毫秒
	Ctime int64
	//更新时间 毫秒
	Utime int64
}
