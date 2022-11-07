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

type ListRequest struct {
	Filters  Filter
	Page     uint
	PageSize uint
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

type ListRepo interface {
	ListClients(filter Filter, pageSize, page uint) ([]Client, *int64, error)
}

type ListUseCase interface {
	List(req ListRequest) (*ListResponse, error)
}

type listUseCase struct {
	logger      Logger
	repo        ListRepo
	maxPageSize uint
}

func NewListUseCase(logger Logger, repo ListRepo, maxPageSize uint) ListUseCase {
	return &listUseCase{logger, repo, maxPageSize}
}

func (s *listUseCase) List(req ListRequest) (*ListResponse, error) {
	if req.PageSize > s.maxPageSize {
		req.PageSize = s.maxPageSize
	}
	clients, total, err := s.repo.ListClients(req.Filters, req.PageSize, req.Page)
	if err != nil {
		s.logger.Err("Error getting clients", err.Error())
		return nil, errors.New(fmt.Sprintf("could not get clients: %s", err))
	}
	return &ListResponse{req.Page, *total, clients}, nil
}
