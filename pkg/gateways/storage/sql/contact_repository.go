package sql

import (
	"gorm.io/gorm"
)

type contactTable struct {
	gorm.Model
	name  string
	phone string
}

type ContactRepo struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepo {
	db.AutoMigrate(&contactTable{})
	return &ContactRepo{db}
}

func (cr *ContactRepo) Delete() error {
	panic("not implemented") // TODO: Implement
}

func (cr *ContactRepo) Add(name, phone string) error {
	cr.db.Create(&contactTable{
		name:  name,
		phone: phone,
	})
	return nil
}
