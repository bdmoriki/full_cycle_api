package database

import (
	"github.com/bdmoriki/full_cycle_api/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetDBUser() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.User{})
	return db, err
}

func GetDBProduct() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.Product{})
	return db, err
}
