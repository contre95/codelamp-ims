package clients

import (
	"errors"
	"fmt"
	"time"
)

type AddClientResponse struct {
	ID ClientID
}

type AddClientRequest struct {
	Name          string
	AdmissionDate string
	FinishDate    string
	Website       string
	Country       string
	Tag           string
}

type AddClientRepo interface {
	AddClient(c Client) (*ClientID, error)
}

type AddingService interface {
	Add(req AddClientRequest) (*AddClientResponse, error)
}

type addingService struct {
	logger Logger
	repo   Repo
}

func NewAddingService(logger Logger, repo Repo) AddingService {
	return &service{logger, repo}
}

func (s *service) Add(req AddClientRequest) (*AddClientResponse, error) {
	s.logger.Info("Creating client: %s ", req.Name)
	newClient, err := parseAddClientRequest(req)
	if err != nil {
		return nil, errors.New("Client data is not valid")
	}
	cid, err := s.repo.AddClient(*newClient)
	if err != nil {
		s.logger.Err("Error creating client: %v", err)
		return nil, errors.New(fmt.Sprintf("%v", err))
	}
	return &AddClientResponse{ID: *cid}, nil
}

func parseAddClientRequest(req AddClientRequest) (*Client, error) {
	var err error
	var admissionDate, finishDate time.Time
	admissionDate, err = time.Parse(time.RFC3339, req.AdmissionDate)
	if req.FinishDate != "" {
		finishDate, err = time.Parse(time.RFC3339, req.FinishDate)
	}
	if err != nil {
		return nil, err
	}
	client := &Client{
		Name:          req.Name,
		AdmissionDate: admissionDate,
		FinishDate:    finishDate,
		Website:       req.Website,
		Country:       req.Country,
		Tag:           req.Tag,
	}
	err = client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}
