package repository

import (
	"context"

	"wp-demo/pkg/domain/model"
	"wp-demo/pkg/domain/repository"

	"gorm.io/gorm"
)

func NewArticleRepository(db *gorm.DB) repository.ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

var _ repository.ArticleRepository = &articleRepository{}

type articleRepository struct {
	db *gorm.DB
}


func (a *articleRepository) Delete(ctx context.Context, id uint) error {
	_, err := gorm.G[model.Article](a.db).Where("id = ?", id).Delete(ctx)
	return err
}

func (a *articleRepository) Get(ctx context.Context, id uint) (*model.Article, error) {
	article, err := gorm.G[model.Article](a.db).Where("id = ?", id).First(ctx)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

func (a *articleRepository) Create(ctx context.Context, article model.Article) error {
	return gorm.G[model.Article](a.db).Create(ctx, &article)
}

func (a *articleRepository) List(ctx context.Context, page, pageSize int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	db := a.db.WithContext(ctx).Model(&model.Article{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func (a *articleRepository) Update(ctx context.Context, article model.Article) error {
	return a.db.WithContext(ctx).Save(&article).Error
}

func (a *articleRepository) ListByAuthor(ctx context.Context, author string, page, pageSize int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	db := a.db.WithContext(ctx).Model(&model.Article{}).Where("author = ?", author)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}