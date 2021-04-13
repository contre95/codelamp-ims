package contact_repo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ContactRepo struct {
	gorm.Model
	Name  string
	Phone string
}
