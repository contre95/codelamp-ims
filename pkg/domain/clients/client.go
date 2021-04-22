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
	AdmissionDate time.Time
	FinishDate    time.Time
	Website       string
	Country       string
	Tag           string
	contacts      []contacts.ContactID
	projects      []Project
}

func (c *Client) ValidateName() error {
	if len(c.Name) > 1 {
		return errors.New("Name too short")
	}
	return nil
}
