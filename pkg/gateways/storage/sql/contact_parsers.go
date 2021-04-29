package sql

import (
	"codelamp-ims/pkg/domain/contacts"
)

func parseDomainContacts(dcs []contacts.Contact) []Contact {
	var contacts []Contact
	for _, dc := range dcs {
		contacts = append(contacts, *parseDomainContact(dc))
	}
	return contacts
}
func parseDomainContact(c contacts.Contact) *Contact {
	return &Contact{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Phone:     c.Phone,
		Comments:  c.Comments,
	}
}

func parseDBContacts(dbcs []Contact) []contacts.Contact {
	var domainContacts []contacts.Contact
	for _, c := range dbcs {
		domainContacts = append(domainContacts, *parseDBContact(c))
	}
	return domainContacts
}

func parseDBContact(dbc Contact) *contacts.Contact {
	return &contacts.Contact{
		ID:        contacts.ContactID(dbc.ID),
		FirstName: dbc.FirstName,
		LastName:  dbc.LastName,
		Phone:     dbc.Phone,
		Comments:  dbc.Comments,
	}
}
