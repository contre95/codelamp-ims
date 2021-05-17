package http

import (
	"codelamp-ims/pkg/domain/clients"

	fiber "github.com/gofiber/fiber/v2"
)

func createClient(s *clients.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("pong")
	}
}
