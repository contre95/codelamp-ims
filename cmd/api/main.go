package main

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/health"
	"codelamp-ims/pkg/gateways/logger"
	"codelamp-ims/pkg/gateways/storage/sql"
	"codelamp-ims/pkg/presenters/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("db/ims.db"), &gorm.Config{})
	storageService := sql.NewStorage(db)
	storageService.Migrate()

	clientLogger := logger.NewSTDLogger("CLIENTS", logger.VIOLET)
	healthLogger := logger.NewSTDLogger("HEALTH", logger.GREEN2)

	healthService := health.NewService(healthLogger)
	addClientSrv := clients.NewAddService(clientLogger, storageService)
	updateClientSrv := clients.NewUpdateService(clientLogger, storageService)
	listClientSrv := clients.NewListService(clientLogger, storageService)

	fiberApp := fiber.New()
	http.MapRoutes(fiberApp, &healthService, &addClientSrv)
	fiberApp.Listen(":3000")
}
