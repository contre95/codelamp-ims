package clients

import (
	"errors"
	"time"
)

type ProjectID = uint
type ProjectType string
type ProjectState string

type Project struct {
	ID            ProjectID
	Name          string
	StartDate     time.Time
	FinishDate    time.Time
	Website       string
	GitRepository string
	Type          ProjectType
	State         ProjectState
	Tag           string
}

func (p *Project) ValidateName() error {
	if len(p.Name) > 1 {
		return errors.New("Name too short")
	}
	return nil
}

func (p *Project) ValidateDates() error {
	if len(p.Name) > 1 {
		return errors.New("Name too short")
	}
	return nil
}
