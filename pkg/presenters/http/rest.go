package http

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/health"

	fiber "github.com/gofiber/fiber/v2"
)

// MapRoutes is where http REST routes are mapped to functions
func MapRoutes(fi *fiber.App, he *health.Service, c *clients.Service) {
	fi.Get("/ping", ping(*he)) // /api/v1/ping
	api := fi.Group("/api")    // /api
	v1 := api.Group("/v1")     // /api/v1
	v1.Post("/clients", createClient(*&c.AddUseCase))
	v1.Get("/clients/:id", getClient(*&c.GetUseCase))
}

func ping(h health.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(h.Ping())
	}
}
