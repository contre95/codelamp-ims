package sql

import "gorm.io/gorm"

type projectTable struct {
	gorm.Model
	name  string
	phone string
}

func (r *SQLStorage) DeleteProject() error {
	panic("not implemented") // TODO: Implement
}

func (r *SQLStorage) AddProject(name, phone string) error {
	r.db.Create(&projectTable{
		name:  name,
		phone: phone,
	})
	return nil

