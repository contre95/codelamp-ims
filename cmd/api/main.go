package main

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/contacts"
	"codelamp-ims/pkg/domain/health"
	"codelamp-ims/pkg/gateways/storage/sql"
	"codelamp-ims/pkg/presenters/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, _ := gorm.Open(sqlite.Open("ims.db"), &gorm.Config{})
	storageService := sql.NewStorage(db)
	storageService.Migrate()

	contactService := contacts.NewService(storageService)
	clientsService := clients.NewService(storageService)
	healthService := health.NewService()
	fiberApp := fiber.New()
	http.MapRoutes(fiberApp, &clientsService, &contactService, &healthService)
	fiberApp.Listen(":3000")
	return
}
