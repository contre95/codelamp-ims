package main

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/contacts"
	"codelamp-ims/pkg/gateways/storage/sql"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hi")
	db, _ := gorm.Open(sqlite.Open("db/ims.db"), &gorm.Config{})
	sqlStorage := sql.NewStorage(db)
	sqlStorage.Migrate()
	fmt.Println("Database migrated")

	contactService := contacts.NewService(sqlStorage)
	clientsService := clients.NewService(sqlStorage)
	fmt.Println("Services instanciated")
	for _, contact := range SampleContacts {
		log.Printf("Adding contact : %s \n", contact.FirstName+" "+contact.LastName)
		contactService.Create(contact)
	}
	for _, client := range SampleClients {
		log.Printf("Adding clients: %s \n", client.Name)
		clientsService.Create(client)
	}
}
