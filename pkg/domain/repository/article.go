package repository

import (
	"context"

	"wp-demo/pkg/domain/model"
)

type ArticleRepository interface {
	Create(context.Context, model.Article) error
	Get(context.Context, uint) (*model.Article, error)
	Delete(context.Context, uint) error
	List(ctx context.Context, page, pageSize int) ([]model.Article, int64, error)
	Update(ctx context.Context, article model.Article) error
	ListByAuthor(ctx context.Context, author string, page, pageSize int) ([]model.Article, int64, error)
}