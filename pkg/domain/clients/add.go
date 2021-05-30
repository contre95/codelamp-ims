package clients

import (
	"errors"
	"fmt"
	"time"
)

type AddResponse struct {
	ID ClientID
}

type AddRequest struct {
	Name          string
	AdmissionDate string
	FinishDate    string
	Website       string
	Country       string
	Tag           string
}

type AddRepo interface {
	AddClient(c Client) (*ClientID, error)
}

type AddUseCase interface {
	Add(req AddRequest) (*AddResponse, error)
}

type addUseCase struct {
	logger Logger
	repo   AddRepo
}

func NewAddUseCase(logger Logger, repo AddRepo) AddUseCase {
	return &addUseCase{logger, repo}
}

func (s *addUseCase) Add(req AddRequest) (*AddResponse, error) {
	s.logger.Info("Creating client: %s ", req.Name)
	newClient, err := parseAddRequest(req)
	if err != nil {
		return nil, errors.New("Client data is not valid")
	}
	cid, err := s.repo.AddClient(*newClient)
	if err != nil {
		s.logger.Err("Error creating client: %v", err)
		return nil, errors.New(fmt.Sprintf("%v", err))
	}
	return &AddResponse{ID: *cid}, nil
}

func parseAddRequest(req AddRequest) (*Client, error) {
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
