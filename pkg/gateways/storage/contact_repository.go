package contact_repo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ContactModel struct {
	gorm.Model
	Name  string
	Phone string
}
type ContactRepo struct {
	db *gorm.DB
}

func NewContactRepository() ContactRepo {
	db, err := gorm.Open(sqlite.Open("ims.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&ContactModel{})
	return ContactRepo{db}
}

func (cr *ContactRepo) Add(name, phone string) {
	cr.db.Create(&ContactModel{
		Name:  name,
		Phone: phone,
	})
}
