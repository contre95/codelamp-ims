package clients

import (
	"codelamp-ims/pkg/domain/contacts"
	"errors"
	"time"
)

type ClientID uint
type ClientType string
type ClientState string

type Client struct {
	ID            ClientID
	Name          string
	FinishDate    time.Time
	AdmissionDate time.Time
	Website       string
	Country       string
	Tag           string
	Contacts      []contacts.ContactID
	Projects      []Project
}

func (c *Client) ValidateName() error {
	if len(c.Name) > 1 {
		return errors.New("Name too short")
	}
	return nil
}
