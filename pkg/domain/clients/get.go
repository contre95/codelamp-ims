package clients

import (
	"errors"
	"fmt"
)

type GetResponse struct {
	Client Client
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
		s.logger.Err("Error getting client with ID: %v", err)
		return nil, errors.New(fmt.Sprint("Couldn't get client data: ", err))
	}
	resp := &GetResponse{*client}
	return resp, nil
}
