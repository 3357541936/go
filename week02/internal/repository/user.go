package repository

import (
	"context"
	"week02.com/internal/domain"
	"week02.com/internal/repository/dao"
)

var ErrDuplicateEmail = dao.ErrDuplicateEmail
var ErrUserNotFound = dao.ErrRecordNotFound

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(ud *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: ud,
	}
}

func (repo *UserRepository) Create(context context.Context, u domain.User) error {
	return repo.dao.Insert(context, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (repo *UserRepository) FindByEmail(context context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(context, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *UserRepository) EditProfile(context context.Context, user domain.User) error {
	return repo.dao.UpdateById(context, dao.User{
		Id:          user.Id,
		Name:        user.Name,
		Birth:       user.Birth,
		Description: user.Description,
	})
}

func (repo *UserRepository) UpdateProfile(context context.Context, user domain.User) error {
	return repo.dao.UpdateById(context, dao.User{
		Id:          user.Id,
		Name:        user.Name,
		Birth:       user.Birth,
		Description: user.Description,
	})
}

func (repo *UserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:          u.Id,
		Email:       u.Email,
		Password:    u.Password,
		Name:        u.Name,
		Birth:       u.Birth,
		Description: u.Description,
	}
}
