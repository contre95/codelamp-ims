package sql

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/contacts"
)

//func (sql *SQLStorage) parseDomainProject(p clients.Project) *Project {
//var contacts []Contact
//for _, ci := range p.Contacts {
//contact, err := sql.GetContact(ci)
//if err != nil {
//log.Println("Could not get contacts from database")
//return nil
//}
//contacts = append(contacts, contact)
//}
//return &Project{
//Name:          p.Name,
//StartDate:     p.StartDate,
//FinishDate:    p.FinishDate,
//Website:       p.Website,
//GitRepository: p.GitRepository,
//Type:          string(p.Type),
//State:         string(p.State),
////Contacts:      []uint,  // Acá me conviene llamar al dominio de contacts o llamo a este mismo repo??
//// TODO: Decide
//}
//}

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
