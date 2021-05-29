package clients

import (
	"errors"
	"fmt"
)

type ListResponse struct {
	Page      uint
	PageTotal int64
	Clients   []Client
}

type Filter struct {
	Name          string
	AdmissionDate string
	FinishDate    string
	Website       string
	Country       string
	Tag           string
	OrderBy       string
}

type ListRequest struct {
	Filters  Filter
	Page     uint
	PageSize uint
}

type ListRepo interface {
	ListClients(filter Filter, pageSize, page uint) ([]Client, int64, error)
}

type ListService interface {
	List(req ListRequest) (*ListResponse, error)
}

type listService struct {
	logger Logger
	repo   ListRepo
}

func NewListService(logger Logger, repo ListRepo) ListService {
	return &listService{logger, repo}
}

func (s *listService) List(req ListRequest) (*ListResponse, error) {
	clients, total, err := s.repo.ListClients(req.Filters, req.PageSize, req.Page)
	if err != nil {
		s.logger.Err("Error getting clients", err)
		return nil, errors.New(fmt.Sprintf("could not get clients: %s", err))
	}
	return &ListResponse{req.Page, total, clients}, nil
}
