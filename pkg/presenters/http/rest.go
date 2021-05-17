package http

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/domain/contacts"
	"codelamp-ims/pkg/domain/health"

	fiber "github.com/gofiber/fiber/v2"
)

// MapRoutes is where http REST routes are mapped to functions
func MapRoutes(fi *fiber.App, cl *clients.Service, co *contacts.Service, he *health.Service) {
	api := fi.Group("/api")   // /api
	v1 := api.Group("/v1")    // /api/v1
	v1.Get("/ping", ping(he)) // /api/v1/ping
}

func ping(h *health.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("pong")
	}
}
