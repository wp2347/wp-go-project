package db

import (
	"wp-demo/pkg/domain/model"

	// "gorm.io/driver/sqlite"
	sqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := Migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.Migrator().AutoMigrate(
		&model.User{},
		&model.Article{},
	)
}
