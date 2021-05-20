package clients

import (
	"codelamp-ims/pkg/domain/contacts"
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type ClientID = uint
type ClientType string
type ClientState string

type Client struct {
	ID            ClientID
	Name          string    `validate:"required,min=3,max=32"`
	AdmissionDate time.Time `validate:"ltcsfield=InnerStructField.FinishDate"`
	FinishDate    time.Time `validate:"gtcsfield=InnerStructField.AdmissionDate"`
	Website       string    `validate:"url"`
	Country       string    `validate:"required,min=3,max=32"`
	Tag           string    `validate:"kebabCase,startswith=sys-|web-|imp-"`
	Contacts      []contacts.ContactID
	Projects      []Project
}

func (c *Client) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		return err
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
