package sql

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/contacts"
	"time"

	"gorm.io/gorm"
)

type ClientID uint
type ClientType string
type ClientState string

type Client struct {
	Name          string
	AdmissionDate time.Time
	FinishDate    time.Time
	Website       string
	Country       string
	Tag           string
	contacts      []contacts.ContactID
	projects      []Project
}

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

func (sql *SQLStorage) AddClient(c clients.Client) (*clients.ClientID, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) ListClients() ([]clients.Client, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) GetClient(id clients.ClientID) (*clients.Client, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) UpdateClient(c clients.Client) error {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) DeleteClient(id clients.ClientID) (*clients.Client, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) AddProject(cid clients.ClientID, p clients.Project) (*clients.ProjectID, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) GetProjects(cid clients.ClientID) ([]clients.Project, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) GetProject(cid clients.ClientID, pid clients.ProjectID) (*clients.Project, error) {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) UpdateProjct(p clients.Project) error {
	panic("not implemented") // TODO: Implement
}

func (sql *SQLStorage) DeleteProject(pid clients.ProjectID) (*clients.Project, error) {
	panic("not implemented") // TODO: Implement
}
