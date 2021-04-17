package sql

import "gorm.io/gorm"

type contactTable struct {
	gorm.Model
	name  string
	phone string
}

func (r *SQLStorage) DeleteContact() error {
	panic("not implemented") // TODO: Implement
}

func (r *SQLStorage) AddContact(name, phone string) error {
	r.db.Create(&contactTable{
		name:  name,
		phone: phone,
	})
	return nil
}
