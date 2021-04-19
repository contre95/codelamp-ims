package main

import (
	"codelamp-ims/pkg/domain/contacts"
	"codelamp-ims/pkg/gateways/storage/sql"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("ims.db"), &gorm.Config{})
	sqlStorage := sql.NewStorage(db)
	sqlStorage.Migrate()
	contactService := contacts.NewService(sqlStorage)

	for _, c := range SampleContacts {
		log.Printf("Adding contact : %s \n", c.FirstName+" "+c.LastName)
		contactService.Create(c)
	}
}
