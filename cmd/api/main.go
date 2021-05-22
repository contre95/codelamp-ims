package main

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/contacts"
	"codelamp-ims/pkg/domain/health"
	"codelamp-ims/pkg/gateways/logger"
	"codelamp-ims/pkg/gateways/storage/sql"
	"codelamp-ims/pkg/presenters/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("ims.db"), &gorm.Config{})
	storageService := sql.NewStorage(db)
	storageService.Migrate()

	healthLogger := logger.NewSTDLogger("HEALTH")

	serviceLogger := logger.NewSTDLogger("DOMAIN")

	contactService := contacts.NewService(serviceLogger, storageService)
	clientsService := clients.NewService(serviceLogger, storageService)
	healthService := health.NewService(healthLogger)

	fiberApp := fiber.New()
	http.MapRoutes(fiberApp, &clientsService, &contactService, &healthService)
	fiberApp.Listen(":3000")
}
