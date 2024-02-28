package service

import "proxy/pkg/repository"

type Service struct {
	Request  *RequestService
	Response *ResponseService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Request:  NewRequestService(*repo.Request),
		Response: NewResponseService(*repo.Response),
	}
}