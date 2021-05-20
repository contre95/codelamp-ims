package http

import (
	"codelamp-ims/pkg/domain/clients"

	fiber "github.com/gofiber/fiber/v2"
)

type Client struct {
	//ID   string
	Name          string `json:"name"`
	AdmissionDate string `json:"admission_date"`
	FinishDate    string `json:"finish_date"`
	Country       string `json:"country"`
	Website       string `json:"website"`
	Tag           string `json:"tag"`
	Contacts      []uint `json:"contacts"`
	//Projects      []Project
}

func createClient(s *clients.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		return c.SendString("pong")
	}
}
