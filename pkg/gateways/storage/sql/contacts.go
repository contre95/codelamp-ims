package sql

func (r *Repository) DeleteContact() error {
	panic("not implemented") // TODO: Implement
}

func (r *Repository) AddContact(name, phone string) error {
	r.db.Create(&contactTable{
		name:  name,
		phone: phone,
	})
	return nil
}
