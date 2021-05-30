package http

import (
	"codelamp-ims/pkg/domain/clients"
	"fmt"
	"net/http"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

func createClient(s clients.AddUseCase) func(*fiber.Ctx) error {
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
		newClientRequest := &clients.AddRequest{
			Name:          newClientData.Name,
			AdmissionDate: newClientData.AdmissionDate,
			FinishDate:    newClientData.FinishDate,
			Website:       newClientData.Website,
			Country:       newClientData.Country,
			Tag:           newClientData.Tag,
		}
		addClientResponse, err := s.Add(*newClientRequest)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success":   false,
				"client_id": nil,
				"err":       fmt.Sprintf("%v", err),
			})
		}
		return c.Status(http.StatusAccepted).JSON(&fiber.Map{
			"success":   true,
			"client_id": addClientResponse.ID,
			"err":       nil,
		})
	}
}

func getClient(s clients.GetUseCase) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, uriError := strconv.Atoi(c.Params("id"))
		if uriError != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"client":  nil,
				"err":     fmt.Sprintf("%v", uriError),
			})
		}
		getRequest := &clients.GetRequest{ID: uint(id)}
		resp, err := s.Get(*getRequest)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"client":  nil,
				"err":     fmt.Sprintf("%v", err),
			})
		}
		return c.Status(http.StatusAccepted).JSON(&fiber.Map{
			"success": true,
			"client":  resp,
			"err":     nil,
		})

	}
}

//func sampleHanlder(s clients.Service) func(*fiber.Ctx) error {
//return func(c *fiber.Ctx) error {
//return nil
//}
//}
