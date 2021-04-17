package sql

import (
	"gorm.io/gorm"
)

type SQLStorage struct {
	db *gorm.DB
}

func NewSQLRepository(db *gorm.DB) *SQLStorage {
	db.AutoMigrate(&contactTable{})
	return &SQLStorage{db}
}
