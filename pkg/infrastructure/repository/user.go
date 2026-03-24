package repository

import (
	"context"

	"wp-demo/pkg/domain/model"
	"wp-demo/pkg/domain/repository"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

var _ repository.UserRepository = &userRepository{}

type userRepository struct {
	db *gorm.DB
}

// Create implements repository.UserRepository.
func (u *userRepository) Create(ctx context.Context, user model.User) error {
	return gorm.G[model.User](u.db).Create(ctx, &user)
}

// Get implements repository.UserRepository.
func (u *userRepository) Get(ctx context.Context, username string) (*model.User, error) {
	user, err := gorm.G[model.User](u.db).Where("user_name = ?", username).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
