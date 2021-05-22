package http

import (
	"codelamp-ims/pkg/domain/clients"
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

func createClient(s clients.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		newClientData := Client{}
		err := c.BodyParser(&newClientData)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success":   false,
				"client_id": nil,
				"err":       fmt.Sprintf("%v", err),
			})
		}
		newClient, err := parseJSONClient(newClientData)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success":   false,
				"client_id": nil,
				"err":       fmt.Sprintf("%v", err),
			})
		}
		id, err := s.Create(*newClient)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success":   false,
				"client_id": nil,
				"err":       fmt.Sprintf("%v", err),
			})
		}
		return c.Status(http.StatusAccepted).JSON(&fiber.Map{
			"success":   true,
			"client_id": id,
			"err":       nil,
		})
	}
}

func getClients(s clients.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		panic("Implement me")
	}
}
func sampleHanlder(s clients.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
