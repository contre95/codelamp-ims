package projects

import (
	"codelamp-ims/pkg/domain/contacts"
	"errors"
	"log"
)

type Repo interface {
	AddProject(c Project) error
	ListProjects() ([]Project, error)
	GetProject(id ProjectID) (*Project, error)
	UpdateContact(c Project) error
	DeleteProject(id ProjectID) (*Project, error)
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
	err := s.repo.AddProject(p)
	if err != nil {
		log.Printf("Could not create project %s \n", p.Name)
		return errors.New("Couldn't create project.")
	}
	return nil
}

func (s *service) Delete(id ProjectID) error {
	p, err := s.repo.DeleteProject(id)
	if err != nil {
		log.Printf("Could not delete project \n")
		return errors.New("Couldn't delete project.")
	}
	log.Printf("Project %s was deleted succesfully", p.Name) // Deletion along with contacts?
	return nil
}

func (s *service) ListProjects() ([]Project, error) {
	panic("not implemented") // TODO: Implement
}

// Should this be part of the project domain or be treated as separate a separate entity ?
func (s *service) AddContactsToProjects(pid ProjectID, cids []contacts.ContactID) {
	panic("not implemented") // TODO: Implement
}
