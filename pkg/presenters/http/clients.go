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
			fmt.Println(err)
			c.Status(http.StatusNotAcceptable)
			return c.JSON(err)
		}
		newClient, err := parseJSONClient(newClientData)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusNotAcceptable)
			return c.JSON(err)
		}
		id, err := s.Create(*newClient)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}
		fmt.Println(err)
		c.Status(http.StatusAccepted)
		return c.JSON(id)
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
