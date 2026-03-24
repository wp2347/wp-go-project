package repository

import (
	"context"

	"wp-demo/pkg/domain/model"
)

type UserRepository interface {
	Create(context.Context, model.User) error
	Get(context.Context, string) (*model.User, error)
}
