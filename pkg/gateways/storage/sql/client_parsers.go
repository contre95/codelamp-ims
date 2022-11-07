package sql

import (
	"codelamp-ims/pkg/domain/clients"
)

func parseDBProjects(dbps []Project) []clients.Project {
	var projects []clients.Project
	for _, dbp := range dbps {
		projects = append(projects, *parseDBProject(dbp))
	}
	return projects
}
func parseDBProject(dbp Project) *clients.Project {
	return &clients.Project{
		ID:            clients.ProjectID(dbp.ID),
		Name:          dbp.Name,
		StartDate:     dbp.StartDate,
		FinishDate:    dbp.FinishDate,
		Website:       dbp.Website,
		GitRepository: dbp.GitRepository,
		Type:          clients.ProjectType(dbp.Type),
		Tag:           dbp.Tag,
		State:         clients.ProjectState(dbp.State),
	}
}

func parseDomainClient(c clients.Client) *Client {
	var dbProjects []Project
	for _, p := range c.Projects {
		dbProjects = append(dbProjects, *parseDomainProject(p))
	}

	return &Client{
		Name:          c.Name,
		FinishDate:    c.FinishDate,
		AdmissionDate: c.AdmissionDate,
		Website:       c.Website,
		Country:       c.Country,
		Tag:           c.Tag,
		Projects:      dbProjects,
	}
}

func parseDBClients(dbcs []Client) []clients.Client {
	var clients []clients.Client
	for _, dbc := range dbcs {
		clients = append(clients, *parseDBClient(dbc))
	}
	return clients
}

func parseDBClient(dbc Client) *clients.Client {
	var projects []clients.Project
	for _, dbp := range dbc.Projects {
		projects = append(projects, *parseDBProject(dbp))
	}

	return &clients.Client{
		ID:            clients.ClientID(dbc.ID),
		Name:          dbc.Name,
		AdmissionDate: dbc.AdmissionDate,
		FinishDate:    dbc.FinishDate,
		Website:       dbc.Website,
		Country:       dbc.Country,
		Tag:           dbc.Tag,
		Projects:      projects,
	}
}

func parseDomainProject(p clients.Project) *Project {
	// I'm not handling contacts retrieval here
	return &Project{
		Name:          p.Name,
		StartDate:     p.StartDate,
		FinishDate:    p.FinishDate,
		Website:       p.Website,
		GitRepository: p.GitRepository,
		Type:          string(p.Type),
		State:         string(p.State),
		Tag:           p.Tag,
	}
}
