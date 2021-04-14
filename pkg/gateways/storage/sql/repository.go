package sql

import (
	"gorm.io/gorm"
)

type contactTable struct {
	gorm.Model
	name  string
	phone string
}

type Repository struct {
	db *gorm.DB
}

func NewSQLRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&contactTable{})
	return &Repository{db}
}

