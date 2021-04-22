package clients

import (
	"errors"
	"fmt"
)

type Repo interface {
	AddClient(c Client) (*ClientID, error)
	ListClients() ([]Client, error)
	GetClient(id ClientID) (*Client, error)
	UpdateClient(c Client) error
	DeleteClient(id ClientID) (*Client, error)

	AddProject(cid ClientID, p Project) (ProjectID, error)
	GetProjects(cid ClientID) ([]Project, error)
	GetProject(cid ClientID) (*Project, error)
	UpdateProjct(p Project) error
	DeleteProject(pid ProjectID) (*Project, error)
}

type Service interface {
	Create(client Client) (*ClientID, error)
	Delete(id ClientID) (*Client, error)
	UpdateInfo(client Client) error
	Get(cid ClientID) (*Client, error)
	List() ([]Client, error)

	AddProjects(cid ClientID, p []Project) error
	CreateProject(cid ClientID, p Project) error
	GetProjects(cid ClientID) ([]Project, error)
	RemoveProjects(cid ClientID) ([]Project, error)
	RemoveProject(cid ClientID, pid ProjectID) (Project, error)
	UpdateProject(cid ClientID, p Project) error
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo}
}

func (s *service) Create(client Client) (*ClientID, error) {
	cid, err := s.repo.AddClient(client)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not create client %s", client.Name))
	}
	errp := s.AddProjects(*cid, client.projects)
	if errp != nil {
		return nil, errors.New(fmt.Sprintf("Could not add projects to client %s", client.Name))
	}
	return cid, nil
}

func (s *service) Get(cid ClientID) (*Client, error) {
	c, err := s.repo.GetClient(cid)
	if err != nil {
		return nil, errors.New("Could not get requested client")
	}
	return c, nil
}

func (s *service) Delete(cid ClientID) (*Client, error) {
	_, errc := s.RemoveProjects(cid)
	if errc != nil {
		return nil, errors.New("Could not clear projects from clients")
	}
	c, err := s.repo.DeleteClient(cid)
	if err != nil {
		return nil, errors.New("Could not delete client")
	}
	return c, nil
}

func (s *service) UpdateInfo(client Client) error {
	err := s.repo.UpdateClient(client)
	if err != nil {
		return errors.New("Could not update client information.")
	}
	// TODO: No sé si está tan bueno updatear los projects | hacer un diff con los nuevos/existentes | borrar todos y agregar los nuevos
	//errcp := s.AddProjects(client.ID, client.projects)
	//if errcp != nil {
	//return errors.New(fmt.Sprintf("Could not attach projects to client %s", client.Name))
	//}
	return nil
}

func (s *service) List( /*itermPerPage, page, int*/ ) ([]Client, error) {
	clients, err := s.repo.ListClients()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get clients: %s", err))
	}
	return clients, nil
}

func (s *service) AddProjects(cid ClientID, projects []Project) error {
	for _, p := range projects {
		err := s.CreateProject(cid, p)
		if err != nil {
			return errors.New("Could not create projects")
		}
	}
	return nil
}

func (s *service) CreateProject(cid ClientID, p Project) error {

}

func (s *service) GetProjects(cid ClientID) ([]Project, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) RemoveProject(cid ClientID, pid ProjectID) (Project, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) RemoveProjects(cid ClientID) ([]Project, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) UpdateProject(cid ClientID, p Project) error {
	panic("not implemented") // TODO: Implement
}
