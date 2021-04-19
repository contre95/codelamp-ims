package contacts

import (
	"errors"
	"fmt"
)

type Repo interface {
	AddContact(c Contact) error
	ListContacts() ([]Contact, error)
	GetContact(id ContactID) (*Contact, error)
	UpdateContact(c Contact) error
	DeleteContact(id ContactID) (*Contact, error)
}

type Service interface {
	Create(c Contact) error
	List() ([]Contact, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo}
}

func (s *service) Create(c Contact) error {
	err := s.repo.AddContact(c)
	if err != nil {
		return errors.New(fmt.Sprintf("could not create contact: %s", err))
	}
	return nil
}

func (s *service) List() ([]Contact, error) {
	contacts, err := s.repo.ListContacts()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get contacts: %s", err))
	}
	return contacts, nil
}
