package projects

import (
	"codelamp-ims/pkg/domain/contacts"
	"errors"
	"time"
)

type ProjectID int
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
	contacts      []contacts.Contact // A mi hace ruido esto
}

func (p *Project) ValidateName() error {
	if len(p.Name) > 1 {
		return errors.New("Name too short")
	}
	return nil
}
