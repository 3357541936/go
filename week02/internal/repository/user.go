package repository

import (
	"context"
	"week02.com/internal/domain"
	"week02.com/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (repo *UserRepository) Create(context context.Context, u domain.User) error {
	return repo.dao.Insert(context, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
