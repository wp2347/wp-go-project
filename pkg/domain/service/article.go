package service

import (
	"context"

	"wp-demo/pkg/domain/model"
	"wp-demo/pkg/domain/repository"

	"gorm.io/gorm"
)

func NewArticleService(repo repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

type ArticleService struct {
	repo repository.ArticleRepository
}

func (u *ArticleService) Create(ctx context.Context, title, content, author string) error {
	return u.repo.Create(ctx, model.Article{
		Title:   title,
		Content: content,
		Author:  author,
	})
}

func (u *ArticleService) Get(ctx context.Context, id uint) (*model.Article, error) {
	return u.repo.Get(ctx, id)
}

func (u *ArticleService) Delete(ctx context.Context, id uint) error {
	return u.repo.Delete(ctx, id)
}

func (u *ArticleService) List(ctx context.Context, page, pageSize int) ([]model.Article, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return u.repo.List(ctx, page, pageSize)
}

func (u *ArticleService) Update(ctx context.Context, id uint, title, content string) error {
	article, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if article == nil {
		return gorm.ErrRecordNotFound
	}

	article.Title = title
	article.Content = content
	return u.repo.Update(ctx, *article)
}

func (u *ArticleService) ListByAuthor(ctx context.Context, author string, page, pageSize int) ([]model.Article, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return u.repo.ListByAuthor(ctx, author, page, pageSize)
}