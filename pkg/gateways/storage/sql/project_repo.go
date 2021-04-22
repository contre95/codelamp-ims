package sql

import (
	"time"

	"gorm.io/gorm"
)

type ProjectID int
type ProjectType string
type ProjectState string

type Project struct {
	gorm.Model
	Name          string
	StartDate     time.Time
	FinishDate    time.Time
	Website       string
	GitRepository string
	Type          string
	State         string
	Tag           string
	Contacts      []Contact
}
