package sql

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/contacts"
)

func parseDBProject(dbp Project) *clients.Project {
	var contactsIDs []contacts.ContactID
	for _, cid := range dbp.Contacts {
		contactsIDs = append(contactsIDs, contacts.ContactID(cid.ID))
	}
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
		Contacts:      contactsIDs,
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
		//Contacts:      c.Contacts,     // Am I getting contacts here ? No I'm not
	}
}

func parseDBClient(dbc Client) *clients.Client {
	var contactsIDs []contacts.ContactID
	for _, p := range dbc.Projects {
		for _, cid := range p.Contacts {
			contactsIDs = append(contactsIDs, contacts.ContactID(cid.ID))
		}
	}
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
		Contacts:      contactsIDs, // Client contacts are all contacts withing its projects
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
