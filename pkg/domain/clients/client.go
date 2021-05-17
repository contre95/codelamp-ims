package clients

import (
	"codelamp-ims/pkg/domain/contacts"
	"errors"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type ClientID = uint
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

func isKebabCase(fl validator.FieldLevel) bool {
	kebabCaseRE := regexp.MustCompile("^[a-z]+(-[a-z]+)*$")
	// More CaseRe https://www.socketloop.com/tutorials/golang-detect-pascal-kebab-screaming-snake-and-camel-cases
	field := fl.Field().String()
	if kebabCaseRE.MatchString(field) {
		return true
	}
	return false
}
