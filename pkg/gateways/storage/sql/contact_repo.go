package sql

import (
	"codelamp-ims/pkg/domain/contacts"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	FirstName string
	LastName  string
	Phone     string
	Comments  string
	ProjectID uint
}

func parseDomainContact(c contacts.Contact) *Contact {
	return &Contact{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Phone:     c.Phone,
		Comments:  c.Comments,
	}
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

func (sql *SQLStorage) AddContact(c contacts.Contact) error {
	Contact := parseDomainContact(c)
	result := sql.db.Create(Contact)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sql *SQLStorage) ListContacts() ([]contacts.Contact, error) {
	var dbContacts []Contact
	result := sql.db.Find(&dbContacts)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not fetch contacts: %s \n", result.Error))
	}
	var domainContacts []contacts.Contact
	for _, dbContact := range dbContacts {
		domainContacts = append(domainContacts, *parseDBContact(dbContact))
	}
	return domainContacts, nil
}

func (sql *SQLStorage) GetContact(id contacts.ContactID) (*contacts.Contact, error) {
	var dbContact Contact
	result := sql.db.First(&dbContact, id)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve contact: %s \n", result.Error))
	}
	return parseDBContact(dbContact), nil
}

func (sql *SQLStorage) UpdateContact(c contacts.Contact) error {
	var dbContact Contact
	result := sql.db.First(dbContact, c.ID)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Could not update contact (contact %s not found): %s \n", c.FirstName, result.Error))
	}
	dbContact.FirstName = c.FirstName
	dbContact.LastName = c.LastName
	dbContact.Phone = c.Phone
	dbContact.Comments = c.Comments
	dbContact.UpdatedAt = time.Now()
	err := sql.db.Save(dbContact)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not update contact (%s): %s \n", c.FirstName, result.Error))
	}
	return nil
}

func (sql *SQLStorage) DeleteContact(id contacts.ContactID) (*contacts.Contact, error) {
	var dbContact Contact
	result_get := sql.db.First(&dbContact, id)
	if result_get.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve contact: %s \n", result_get.Error))
	}
	result_delete := sql.db.Delete(&Contact{}, id)
	if result_delete.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not delete contact: %s \n", result_delete.Error))
	}
	return parseDBContact(dbContact), nil
}
