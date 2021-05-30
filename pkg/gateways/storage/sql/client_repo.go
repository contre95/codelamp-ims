package sql

import (
	"codelamp-ims/pkg/domain/clients"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ClientID uint
type ClientType string
type ClientState string

type Client struct {
	gorm.Model
	Name          string `gorm:"unique"`
	FinishDate    time.Time
	AdmissionDate time.Time
	Website       string
	Country       string
	Tag           string
	Projects      []Project
	//Contacts      []contacts.ContactID
}

type ProjectID int
type ProjectType string
type ProjectState string

type Project struct {
	gorm.Model
	Name          string `gorm:"unique"`
	StartDate     time.Time
	FinishDate    time.Time
	Website       string
	GitRepository string
	Type          string
	State         string
	Tag           string
	ClientID      uint
	Contacts      []Contact
}

func (sql *SQLStorage) AddClient(c clients.Client) (*clients.ClientID, error) {
	client := parseDomainClient(c)
	result := sql.db.Create(client)
	if result.Error != nil {
		return nil, result.Error
	}
	return &client.ID, nil
}

func (sql *SQLStorage) ListClients(filter clients.Filter, pageSize, page uint) ([]clients.Client, *int64, error) {
	var clients []Client
	var count int64
	sql.db.Model(&Client{}).Count(&count)
	result := sql.db.Scopes(sql.paginate(pageSize, page)).Find(&clients)
	if result.Error != nil {
		return nil, nil, errors.New(fmt.Sprintf("Could not fetch contacts: %s \n", result.Error))
	}
	return parseDBClients(clients), &count, nil
}

func (sql *SQLStorage) GetClient(id clients.ClientID) (*clients.Client, error) {
	var dbClient Client
	result := sql.db.First(&dbClient, id)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve clients: %s \n", result.Error))
	}
	return parseDBClient(dbClient), nil
}

func (sql *SQLStorage) DeleteClient(id clients.ClientID) (*clients.Client, error) {
	var dbClient Client
	result_get := sql.db.First(&dbClient, id)
	if result_get.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve client: %s \n", result_get.Error))
	}
	result_delete := sql.db.Delete(&Client{}, id)
	if result_delete.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not delete client: %s \n", result_get.Error))
	}
	return parseDBClient(dbClient), nil
}

func (sql *SQLStorage) UpdateClientDetails(c clients.Client) error {
	var dbClient Client
	result := sql.db.First(dbClient, c.ID)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Could not update client (client %s not found): %s \n)", c.Name, result.Error))
	}
	dbClient.Name = c.Name
	dbClient.FinishDate = c.FinishDate
	dbClient.AdmissionDate = c.AdmissionDate
	dbClient.Website = c.Website
	dbClient.Country = c.Country
	dbClient.Tag = c.Tag
	err := sql.db.Save(dbClient)
	if err != nil {
		return errors.New(fmt.Sprintf("Coudl not update client (%s): %v \n", c.Name, err))
	}
	return nil
}

func (sql *SQLStorage) AddProject(cid clients.ClientID, p clients.Project) (*clients.ProjectID, error) {
	var dbClient Client
	result := sql.db.First(&dbClient, cid)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could noot retrieve client for update: %s \n", result.Error))
	}
	// Create project
	project := parseDomainProject(p)
	result_project := sql.db.Create(&project)
	if result_project != nil {
		return nil, errors.New(fmt.Sprintf("Could not add project, to "))
	}
	// end Create project
	dbClient.Projects = append(dbClient.Projects, *project)
	err := sql.db.Save(dbClient)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Coudl not update client (%s): %v \n", dbClient.Name, err))
	}
	return &project.ID, nil
}

func (sql *SQLStorage) GetProjects(cid clients.ClientID) ([]clients.Project, error) {
	var dbClient Client
	result := sql.db.First(dbClient, cid)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Coudl not get client's projects (%s)\n", dbClient.Name))
	}
	return parseDBProjects(dbClient.Projects), nil
}

func (sql *SQLStorage) GetProject(cid clients.ClientID, pid clients.ProjectID) (*clients.Project, error) {
	var dbProjects Project
	result := sql.db.First(dbProjects, pid)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve project: %s \n", result.Error))
	}
	return parseDBProject(dbProjects), nil
}

func (sql *SQLStorage) UpdateProjct(p clients.Project) error {
	var dbProject Project
	result := sql.db.First(dbProject, p.ID)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Could not update client (client %s not found): %s \n)", p.Name, result.Error))
	}
	dbProject.Name = p.Name
	dbProject.StartDate = p.StartDate
	dbProject.FinishDate = p.FinishDate
	dbProject.Website = p.Website
	dbProject.GitRepository = p.GitRepository
	dbProject.Type = string(p.Type)
	dbProject.State = string(p.State)
	dbProject.Tag = p.Tag

	err := sql.db.Save(dbProject)
	if err != nil {
		return errors.New(fmt.Sprintf("Coudl not update client (%s): %v \n", p.Name, err))
	}
	return nil
}

func (sql *SQLStorage) DeleteProject(pid clients.ProjectID) (*clients.Project, error) {
	var dbProject Project
	result_get := sql.db.First(&dbProject, pid)
	if result_get.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve project: %s \n", result_get.Error))
	}
	result_delete := sql.db.Delete(&Project{}, pid)
	if result_delete.Error != nil {
		return nil, errors.New(fmt.Sprintf("Could not delete project: %s \n", result_delete.Error))
	}
	return parseDBProject(dbProject), nil

}
