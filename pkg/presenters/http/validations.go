package http

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Job struct {
	Type   string `validate:"required,min=3,max=32"`
	Salary int    `validate:"required,number"`
}

type User struct {
	Name     string `validate:"required,min=3,max=32"`
	IsActive bool   `validate:"required,eq=True|eq=False"`
	Email    string `validate:"required,email,min=6,max=32"`
	Job      Job    `validate:"dive"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Client struct {
	//ID   string
	Name          string `validate:"required,min=3,max=32"`
	AdmissionDate string `validate:"ltcsfield=InnerStructField.FinishDate"`
	FinishDate    string `validate:"gtcsfield=InnerStructField.AdmissionDate"`
	Country       string `validate:"required,min=3,max=32"`
	Website       string `validate:"url"`
	Tag           string `validate:"kebabCase,startswith=sys-|web-|imp-"`
	Contacts      []uint
	Projects      []Project
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
