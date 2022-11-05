package main

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/health"
	"codelamp-ims/pkg/gateways/logger"
	"codelamp-ims/pkg/gateways/storage/sql"
	"codelamp-ims/pkg/presenters/http"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("db/ims.db"), &gorm.Config{})
	storage := sql.NewStorage(db)
	storage.Migrate()

	clientLogger := logger.NewSTDLogger("CLIENTS", logger.VIOLET)
	healthLogger := logger.NewSTDLogger("HEALTH", logger.GREEN2)

	healthService := health.NewService(healthLogger)

	add := clients.NewAddUseCase(clientLogger, storage)
	list := clients.NewListUseCase(clientLogger, storage, 100)
	get := clients.NewGetUseCase(clientLogger, storage)
	update := clients.NewUpdateUseCase(clientLogger, storage)
	del := clients.NewDeleteUseCase(clientLogger, storage)

	clientService := clients.NewService(add, list, get, update, del)

	switch os.Args[1:][0] {
	case "rest":
		fiberApp := fiber.New()
		http.MapRoutes(fiberApp, &healthService, &clientService)
		fiberApp.Listen(":3000")
	case "graphql":
		graphql.MapRoutes()
		log.Fatal(http.ListenAndServe(":3001", nil))

	}
}
