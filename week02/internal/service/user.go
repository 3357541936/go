package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"week02.com/internal/domain"
	"week02.com/internal/repository"
)

var ErrDuplicateEmail = repository.ErrDuplicateEmail
var ErrInvailidEmailOrPassword = repository.ErrUserNotFound

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (service *UserService) Signup(context context.Context, u domain.User) error {
	password := []byte(u.Password)
	// 密码加密部分,
	n_password, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	u.Password = string(n_password)
	//bcrypt.CompareHashAndPassword()
	return service.repo.Create(context, u)
}

func (service *UserService) Login(context context.Context, email string, password string) (domain.User, error) {
	u, err := service.repo.FindByEmail(context, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvailidEmailOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvailidEmailOrPassword
	}
	return u, nil
}

func (service *UserService) Edit(context context.Context, u domain.User) error {
	return service.repo.EditProfile(context, u)
}

func (service *UserService) Profile(context context.Context, u domain.User) error {
	return service.repo.UpdateProfile(context, u)
}
