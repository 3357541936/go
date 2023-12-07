package service

import (
	"context"
	"week02.com/internal/domain"
	"week02.com/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// service 模块依赖 Domain模块

func (service *UserService) Signup(context context.Context, u domain.User) error {
	return service.repo.Create(context, u)
}
