package model

import (
	"time"	
	"gorm.io/gorm"
)

type Article struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
	Author  string `gorm:"column:author" json:"author"`
	// gorm.Model
}

func (u *Article) Table() string {
	return "article"
}
