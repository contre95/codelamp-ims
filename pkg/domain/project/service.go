package projects

import "codelamp-ims/pkg/domain/contacts"

type Repo interface {
	AddProject(name, phone string) error
	ListProjects() ([]Project, error)
	GetProjects(id int) (Project, error)
	UpdateProject(name, phone string) error
	DeleteProject() error
}

type Service interface {
	Create(name, url string, Projects []contacts.Contact) error
	Delete(id int) error
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo}
}

func (s *service) Create(name string, url string, Projects []contacts.Contact) error {
	panic("not implemented") // TODO: Implement
}

func (s *service) Delete(id int) error {
	panic("not implemented") // TODO: Implement
}
