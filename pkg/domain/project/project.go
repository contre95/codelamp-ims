package projects

import "codelamp-ims/pkg/domain/contacts"

type Project struct {
	ID       int
	Name     string
	Contacts []contacts.Contact
}
