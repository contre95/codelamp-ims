package clients

import (
	"errors"
	"time"
)

type UpdateResponse struct {
	ID ClientID
}

type UpdateRequest struct {
	Name          string
	AdmissionDate string
	FinishDate    string
	Website       string
	Country       string
	Tag           string
}

type UpdateRepo interface {
	UpdateClientDetails(c Client) error
}

type UpdateService interface {
	UpdateDetailsService(req UpdateRequest) (*UpdateResponse, error)
}

type updateService struct {
	logger Logger
	repo   UpdateRepo
}

func NewUpdateService(logger Logger, repo UpdateRepo) UpdateService {
	return &updateService{logger, repo}
}

func (s *updateService) UpdateDetailsService(req UpdateRequest) (*UpdateResponse, error) {
	s.logger.Info("updating client: %s ", req.Name)
	client, err := parseUpdateRequest(req)
	if err != nil {
		return nil, errors.New("Client data is not valid")
	}
	err = s.repo.UpdateClientDetails(*client)
	if err != nil {
		s.logger.Err("Error updating client", err)
		return nil, errors.New("Could not update client information.")
	}
	return &UpdateResponse{}, nil
}

func parseUpdateRequest(req UpdateRequest) (*Client, error) {
	// Here we can enforce which fields can not be updated once the client is created (by not updating them)
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
