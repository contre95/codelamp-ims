package main

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/contacts"
	"codelamp-ims/pkg/gateways/logger"
	"codelamp-ims/pkg/gateways/storage/sql"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hi")
	db, _ := gorm.Open(sqlite.Open("db/ims.db"), &gorm.Config{})
	sqlStorage := sql.NewStorage(db)
	sqlStorage.Migrate()
	fmt.Println("Database migrated")

	contactLogger := logger.NewSTDLogger("CONTACT", logger.BEIGE)
	clientLogger := logger.NewSTDLogger("CLIENTS", logger.VIOLET)

	contactService := contacts.NewService(contactLogger, sqlStorage)
	clientsService := clients.NewService(clientLogger, sqlStorage)
	fmt.Println("Services instanciated")
	for _, contact := range SampleContacts {
		contactService.Create(contact)
	}
	for _, client := range SampleClients {
		clientsService.Create(client)
	}
}
