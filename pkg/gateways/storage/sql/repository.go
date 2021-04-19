package sql

import (
	"gorm.io/gorm"
)

type SQLStorage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *SQLStorage {
	return &SQLStorage{db}
}

func (sql *SQLStorage) Migrate() {
	sql.db.AutoMigrate(&contactDB{})

}
