package clients

import (
	"errors"
	"time"
)

type DeleteResponse struct {
	Name          string
	AdmissionDate time.Time
	FinishDate    time.Time
	Website       string
	Country       string
	Tag           string
	Projects      []Project
}

type DeleteRequest struct {
	ID ClientID
}

type DeleteRepo interface {
	DeleteClient(id ClientID) (*Client, error)
}

type DeleteUseCase interface {
	Delete(req DeleteRequest) (*DeleteResponse, error)
}

type deleteUseCase struct {
	logger Logger
	repo   DeleteRepo
}

func NewDeleteUseCase(logger Logger, repo DeleteRepo) DeleteUseCase {
	return &deleteUseCase{logger, repo}
}

func (s *deleteUseCase) Delete(req DeleteRequest) (*DeleteResponse, error) {
	client, err := s.repo.DeleteClient(req.ID)
	if err != nil {
		s.logger.Err("Error updating client", err.Error())
		return nil, errors.New("Could not Delete client information.")
	}
	resp := &DeleteResponse{
		Name:          client.Name,
		AdmissionDate: client.AdmissionDate,
		FinishDate:    client.FinishDate,
		Website:       client.Website,
		Country:       client.Country,
		Tag:           client.Tag,
		Projects:      client.Projects,
	}
	return resp, nil
}
