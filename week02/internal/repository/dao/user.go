package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound
var ErrDuplicateEmail = errors.New("邮箱(Email)重复")

type UserDAO struct {
	db *gorm.DB
}

// 数据库字段
type User struct {
	Id          int64  `gorm:"primaryKey,autoIncrement"`
	Email       string `gorm:"unique"`
	Password    string
	Name        string
	Birth       int64
	Description string
	CTime       int64
	UTime       int64
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (d *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := d.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (ud *UserDAO) Insert(context context.Context, u User) error {
	u.CTime = time.Now().UnixMilli()
	u.UTime = time.Now().UnixMilli()
	// 根据 SQL 码判定具体错误信息
	err := ud.db.WithContext(context).Create(&u).Error
	if m, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if m.Number == duplicateErr {
			return ErrDuplicateEmail
		}
	}
	return err
}

func (ud *UserDAO) UpdateById(context context.Context, u User) error {
	db_u := User{
		Birth:       u.Birth,
		Name:        u.Name,
		Description: u.Description,
		UTime:       time.Now().UnixMilli(),
	}
	return ud.db.Updates(&User{}).Where("id=?", u.Id).Updates(db_u).Error
}
