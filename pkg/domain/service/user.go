package service

import (
	"context"
	"errors"
	"wp-demo/pkg/domain/model"
	"wp-demo/pkg/domain/repository"
)

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

type UserService struct {
	repo repository.UserRepository
}

func (u *UserService) Register(ctx context.Context, username, password string) error {
	user, err := u.repo.Get(ctx, username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("用户已存在")
	}
	return u.repo.Create(ctx, model.User{
		UserName: username,
		Password: password,
	})
}
