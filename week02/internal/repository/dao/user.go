package dao

import (
	"context"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}

}

func (dao *UserDAO) Insert(context context.Context, u User) error {
	return dao.db.WithContext(context).Create(&u).Error
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	CTime    int64
	UTime    int64
}
