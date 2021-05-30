package clients

import (
	"errors"
	"time"
)

type GetResponse struct {
	Name          string
	AdmissionDate time.Time
	FinishDate    time.Time
	Website       string
	Country       string
	Tag           string
}

type GetRequest struct {
	ID ClientID
}

type GetRepo interface {
	GetClient(id ClientID) (*Client, error)
}

type GetUseCase interface {
	Get(req GetRequest) (*GetResponse, error)
}

type getUseCase struct {
	logger Logger
	repo   GetRepo
}

func NewGetUseCase(logger Logger, repo GetRepo) GetUseCase {
	return &getUseCase{logger, repo}
}

func (s *getUseCase) Get(req GetRequest) (*GetResponse, error) {
	client, err := s.repo.GetClient(req.ID)
	if err != nil {
		s.logger.Err("Error updating client", err)
		return nil, errors.New("Could not get client information.")
	}
	resp := &GetResponse{
		Name:          client.Name,
		AdmissionDate: client.AdmissionDate,
		FinishDate:    client.FinishDate,
		Website:       client.Website,
		Country:       client.Country,
		Tag:           client.Tag,
	}
	return resp, nil
}
