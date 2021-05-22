package contacts

import (
	"errors"
	"fmt"
)

type Logger interface {
	Info(format string, i ...interface{})
	Warn(format string, i ...interface{})
	Err(format string, i ...interface{})
	Debug(format string, i ...interface{})
}

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
	logger Logger
	repo   Repo
}

func NewService(logger Logger, repo Repo) Service {
	return &service{logger, repo}
}

func (s *service) Create(c Contact) error {
	s.logger.Info("Creating contact %s", c.FirstName+" "+c.LastName)
	err := s.repo.AddContact(c)
	if err != nil {
		s.logger.Err("Error creating contact : %v", err)
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
