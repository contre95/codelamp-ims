package projects

import "codelamp-ims/pkg/domain/contacts"

type Repo interface {
	AddProject(name, phone string) error
	ListProjects() ([]Project, error)
	GetProject(id ProjectID) (Project, error)
	UpdateProject(name, phone string) error
	DeleteProject() error
}

type Service interface {
	Create(project Project) error
	Delete(id ProjectID) error
	AddContactsToProjects(pid ProjectID, cids []contacts.ContactID)
	ListProjects() ([]Project, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo}
}

func (s *service) Create(p Project) error {
	panic("not implemented") // TODO: Implement
}

func (s *service) Delete(id ProjectID) error {
	panic("not implemented") // TODO: Implement
}

func (s *service) ListProjects() ([]Project, error) {
	panic("not implemented") // TODO: Implement
}

// Should this be part of the project domain or be treated as separate a separate entity ?
func (s *service) AddContactsToProjects(pid ProjectID, cids []contacts.ContactID) {
	panic("not implemented") // TODO: Implement
}
