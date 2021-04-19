package sql

import (
	"codelamp-ims/pkg/domain/contacts"
	"time"

	"gorm.io/gorm"
)

type ProjectID int
type ProjectType string
type ProjectState string

type projectDB struct {
	gorm.Model
	ID            ProjectID
	Name          string
	StartDate     time.Time
	FinishDate    time.Time
	Website       string
	GitRepository string
	Type          ProjectType
	State         ProjectState
	Tag           string
	Contacts      []contacts.Contact // A mi hace ruido esto
}
